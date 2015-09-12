package desktop

//
// Desktop Folders
//

// user application data folder
func GetAppDataFolder() string {
  return getAppDataFolder();
}

// user home "/home/user"
func GetHomeFolder() string {
  return getHomeFolder()
}

// user my documents "~/Documents"
func GetDocumentsFolder() string {
  return getDocumentsFolder()
}

// user downloads "~/Downloads"
func GetDownloadsFolder() string {
  return getDownloadsFolder()
}

// user desktop "~/Desktop"
func GetDesktopFolder() string {
  return getDesktopFolder()
}

//
// SysTrayIcon or NSStatusBar or Notification Area
//

const (
  MenuItem = 1
  MenuSeparator = 2
  MenuCheckBox = 3
)

type Menu struct {
  Menu []Menu
  Type int
  Enabled bool
  Name string
  Icon []byte
}

type DesktopSysTray struct {
  Listeners map[DesktopSysTrayListener]bool
  Icon []byte
  Title string
  Menu []Menu
  
  // os specific structs
  os interface{}
}

type DesktopSysTrayListener interface {
  MouseLeftClick()

  MouseLeftDoubleClick()

  // We do not handle right clicks, because:
  //
  // 1) Icon is binded to context menu anyway.
  //
  // 2) On Windows if you call showContextMenu from java thread, HMENU bugged
  // and you can't use it.
  //
  // 3) Mac OSX does not support showing context menu programmatically.
}

func DesktopSysTrayNew() *DesktopSysTray {
  return desktopSysTrayNew()
}

func (m *DesktopSysTray) AddListener(l DesktopSysTrayListener) {
  m.Listeners[l] = true;
}

func (m *DesktopSysTray) RemoveListener(l DesktopSysTrayListener) {
  delete(m.Listeners, l);
}

func (m *DesktopSysTray) SetIcon(icon []byte) {
  m.Icon = icon
  update(m)
}

func (m *DesktopSysTray) SetTitle(title string) {
  m.Title = title
  update(m)
}

func (m *DesktopSysTray) Show() {
  show(m)
}

func (m *DesktopSysTray) Update() {
  update(m)
}

func (m *DesktopSysTray) Hide() {
  hide(m)
}

func (m *DesktopSysTray) SetMenu(menu []Menu) {
  m.Menu = menu
  update(m)
}

func (m *DesktopSysTray) Close() {
  close(m)
}
