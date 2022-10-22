package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/x0y14/pm1/command"
	"github.com/x0y14/pm1/password"
)

var (
	errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
)

func (m Model) View() string {
	var s string
	if m.err != nil {
		s += errStyle.Render(fmt.Sprintf("%v", m.err))
		s += "\n"
	}
	s += m.MainView.Render(m)
	s += "\n"
	return s
}

type View struct {
	Action func(m Model) Model
	Render func(m Model) string

	encrypted    []byte
	iv           []byte
	tempHashedMp []byte
}

// FindingEncJson
// 起動されたらまず初めに表示される画面で、
// エクスポートされたファイルを探している間、表示され続けます。
var FindingEncJson = View{
	Action: func(m Model) Model {
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("finding %s\n", exportPath)
	},
}

// CheckEncJson
// ファイルの中身を正常に読み取れるか確認している間、表示され続けます。
var CheckEncJson = View{
	Action: func(m Model) Model {
		encrypted, iv, err := password.Load(exportPath)
		if err != nil {
			// ファイルデータを正常に読み取れませんでした。
			m.err = fmt.Errorf("file data could not be read successfully: %v", err)
			// 終了
			tea.Quit()
		}
		m.MainView = WaitingForToFinishEnteringMasterPasswordForDecrypt
		m.masterPasswordInput.Focus()
		m.MainView.encrypted = encrypted
		m.MainView.iv = iv
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("file checking...\n")
	},
}

// WaitingForToFinishEnteringMasterPasswordForDecrypt
// すでにストレージが存在している場合、
// 復号するためのマスターパスワードを入力してもらうため表示されます。
var WaitingForToFinishEnteringMasterPasswordForDecrypt = View{
	Action: func(m Model) Model {
		// 入力されたデータを取得します
		mp := m.masterPasswordInput.Value()
		// 空白だったらエンターを押されなかったことにします。
		if mp == "" {
			return m
		}

		// 復号
		decryptedStorageBytes, err := password.Decrypt(
			m.MainView.encrypted,
			password.Sha256Hashing(mp),
			m.MainView.iv)
		if err != nil {
			// 復号に失敗したので再度入力してもらいます。
			m.err = fmt.Errorf("decryption failed: %v", err)
			return m
		}

		// 復号できたものをストレージとして読み込みます。
		storage, err := password.LoadStorage(decryptedStorageBytes)
		if err != nil {
			// 復号できたものの、ストレージとして適していないデータでした。
			m.err = fmt.Errorf("could not read as storage: %v", err)
			// 終了します。
			tea.Quit()
		}
		m.err = nil
		m.MainView.iv = nil
		m.MainView.encrypted = nil
		m.masterPasswordInput.Reset()
		m.masterPasswordInput.Blur()
		m.storage = storage
		m.MainView = nextView(m)
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("please entering master password\n%s\n", m.masterPasswordInput.View())
	},
}

// CreatingNewStorageAndVault1
// 初回起動時に表示されます。
// マスターパスワードの入力を要求します。
var CreatingNewStorageAndVault1 = View{
	Action: func(m Model) Model {
		mp := m.masterPasswordInput.Value()
		// 空白だったら再度入力してもらいます。
		if mp == "" {
			return m
		}
		// ５文字未満だったら再度入力してもらいます。
		if len(mp) < 5 {
			m.err = fmt.Errorf("master password entered is too short: at least 5 characters")
			return m
		}
		// 入力クリア
		m.masterPasswordInput.Reset()

		m.MainView = CreatingNewStorageAndVault2
		m.MainView.tempHashedMp = password.Sha256Hashing(mp)
		m.err = nil
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("please enter master password (length: 4 < n)\n%s\n", m.masterPasswordInput.View())
	},
}

// 循環参照エラーを避けるため、代入を遅延します。
var creating1 View

func init() {
	creating1 = CreatingNewStorageAndVault1
}

// CreatingNewStorageAndVault2
// 再度マスターパスワードの入力を要求します。
// 一致しなかった場合は1に戻ります
var CreatingNewStorageAndVault2 = View{
	Action: func(m Model) Model {
		mp := m.masterPasswordInput.Value()
		if string(password.Sha256Hashing(mp)) != string(m.MainView.tempHashedMp) {
			m.err = fmt.Errorf("master passwords dose not match")
			m.masterPasswordInput.Reset()
			m.MainView = creating1
			m.MainView.tempHashedMp = nil
			return m
		}

		// 入力クリア
		m.masterPasswordInput.Reset()
		m.masterPasswordInput.Blur()
		// 個人用保管庫を作成します。
		personalVault := password.NewVault("personal")
		m.storage = password.NewStorage()
		m.storage.Register(personalVault)
		// 作成したストレージをファイル出力して保存します。
		storageBytes, err := password.DumpStorage(m.storage)
		if err != nil {
			// ストレージのダンプに失敗しました
			m.err = fmt.Errorf("dump failed: %v", err)
			// 終了
			tea.Quit()
		}
		// 暗号化します。
		encrypted, iv, err := password.Encrypt(storageBytes, password.Sha256Hashing(mp))
		if err != nil {
			// 暗号化に失敗
			m.err = fmt.Errorf("encrypt failed: %v", err)
			// 終了
			tea.Quit()
		}
		// ファイルへ出力
		err = password.Export(exportPath, encrypted, iv)
		if err != nil {
			// 暗号化したデータのファイル出力に失敗
			m.err = fmt.Errorf("export failed: %v", err)
			// 終了
			tea.Quit()
		}
		m.MainView = nextView(m)
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("please re-enter master password\n%s\n", m.masterPasswordInput.View())
	},
}

func nextView(m Model) View {
	switch m.cmd.Mode {
	case command.MVault:
		switch m.cmd.VaultOpt.Subcommand {
		case "list":
			return vaultListView(m.storage)
		}
	}
	return Success
}

var Success = View{
	Action: func(m Model) Model {
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("succeeded in loading storage.")
	},
}

func ErrorView(err error) View {
	return View{
		Action: func(m Model) Model {
			tea.Quit()
			return m
		},
		Render: func(m Model) string {
			return fmt.Sprintf("error: %v", err)
		},
	}
}
