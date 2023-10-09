import {
  app,
  BrowserWindow,
  ipcMain,
  dialog
} from 'electron';
import {join} from 'path';

function createWindow () {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: join(__dirname, './preload.js')
    }
  })

  let url = "http://127.0.0.1:5174/ui"; // 本地启动的vue项目路径
  win.loadURL(url);
  win.webContents.openDevTools();
}

app.whenReady().then(() => {
  createWindow()

  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow()
    }
  })
})

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})