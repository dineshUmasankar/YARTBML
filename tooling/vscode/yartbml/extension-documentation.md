# YARTBML Extension Documentation

## What's in the folder

- This folder contains all the files necessary to build our extension.
- `package.json` - this is the manifest file in which you declare your language support and define the location of the grammar file that has been copied into your extension.
- `syntaxes/ybml.tmLanguage.json` - this is the Text mate grammar file that is used for tokenization.
- `language-configuration.json` - this is the language configuration, defining the tokens that are used for comments and brackets.

## Developing and testing changes for the extension

- Open the YARTBML project from the `tooling/vscode/yartbml` folder.
- You can launch the extension from the debug toolbar to launch the extension development host.
- You can also reload (`Ctrl+R` or `Cmd+R` on Mac) the VS Code window with your extension to load your changes.

## Install extension from source

### Requirements:
- [Node](https://nodejs.org/en)
- vsce ("Visual Studio Code Extensions") package

To install the vsce, run the following command:
```sh
npm install -g @vscode/vsce
```

- Ensure the YARTBML project is open from the `tooling/vscode/yartbml` folder.
- Open up the terminal and run the following command:
    ```sh
    vsce package
    ```

1. To start using your extension with Visual Studio Code copy the `.vsix` file into the `<user home>/.vscode/extensions` folder and restart VSCode.

or

1. Right-Click the `.vsix` file from the explorer panel in VSCode, and Click `Install from VSIX`.

or

1. Open up the extensions panel (`CTRL + SHIFT + X`) and click the `...` at the upper right
2. Select "Install From VSIX" from the dropdown
3. Traverse the directory until you reach the current folder, where the generated `.vsix` file should be