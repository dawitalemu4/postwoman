## Startup Shortcuts


### Windows

1. Create a new shortcut with the target as the path to your shell of choice by right clicking anywhere on your screen saver and hitting new > shortcut. 

2. Open up the properties of the shortcut and fix the `target` to call the [local startup script](https://github.com/dawitalemu4/postwoman/blob/main/startup.sh) or [docker startup script](https://github.com/dawitalemu4/postwoman/blob/main/.docker-setup/startup.sh) in your postwoman folder and set the `start in` as the path to directory of postwoman.

> my target path: `"C:\Program Files\Git\bin\bash.exe" -i -l -c './startup.sh'"` and my start in path: `D:/developer/postwoman`

3. If you want an icon for your shortcut, hit the change icon button in the properties and use the `favicon.ico` in the views/public folder.

4. Exit the properties and double click the shortcut to start, then from the running shortcut in the taskbar, right click and hit pin to taskbar.

My Final Result:

![Screenshot 2024-04-22 172446](https://github.com/dawitalemu4/postwoman/assets/106638403/a745f410-d117-4d8d-a5a5-02e73a256b6b)


### Mac

1. Create a new shell shortcut in the mac shortcuts app.

2. Copy the [local startup script](https://github.com/dawitalemu4/postwoman/blob/main/startup.sh) or [docker startup script](https://github.com/dawitalemu4/postwoman/blob/main/.docker-setup/startup.sh) into the command field with the shell type being zsh.

3. Click the shortcuts icon on the mac toolbar and click on the shortcut you just created to run postwoman.

4. Hit the square on the shortcut to stop postwoman.

Final Result:


