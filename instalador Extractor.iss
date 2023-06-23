#define MyAppName "Extractor"
#define MyAppVersion "1.0"
#define comentario "Aplicação"
#define folder "Extractor"
#define app "WebVendasProcess"

[Setup]
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
DisableDirPage=yes
DefaultDirName={%HOMEPATH}\{#folder}
DisableProgramGroupPage=yes
DefaultGroupName={#MyAppName}
OutputDir=Instaladores
OutputBaseFilename={#folder}-install
SetupIconFile={#folder}\icon.ico
Compression=lzma
SolidCompression=yes

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"
Name: "brazilianportuguese"; MessagesFile: "compiler:Languages\BrazilianPortuguese.isl"

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"
Name: "brazilianportuguese"; MessagesFile: "compiler:Languages\BrazilianPortuguese.isl"

[Files]
Source: "Template\APP.properties"; DestDir: {%HOMEPATH}\{#folder}; AfterInstall: FileReplaceString(ExpandConstant('{%HOMEPATH}\{#folder}\APP.properties'),'{pasta}', ExpandConstant('{#folder}'), '{jar}', ExpandConstant('{#app}'));  Flags: uninsneveruninstall onlyifdoesntexist ignoreversion
Source: "{#folder}\icon.ico"; DestDir: {%HOMEPATH}\{#folder}; Flags: uninsneveruninstall onlyifdoesntexist ignoreversion
Source: "Template\JVM.properties"; DestDir: {%HOMEPATH}\{#folder}; Flags: uninsneveruninstall onlyifdoesntexist ignoreversion
Source: "Template\MakeManager.jar"; DestDir: {%HOMEPATH}\{#folder}; Flags: uninsneveruninstall onlyifdoesntexist ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Icons]
Name: "{group}\{cm:UninstallProgram,{#MyAppName}}"; Filename: "{uninstallexe}"
;cria no menu iniciar
Name: "{userdesktop}\{#MyAppName}"; Filename: {%HOMEPATH}\{#folder}\MakeManager.jar; Parameters: "t t";  WorkingDir: {%HOMEPATH}\{#folder}; IconFilename: {%HOMEPATH}\{#folder}\icon.ico; Comment: "{#comentario}"

[Run]
Filename: {%HOMEPATH}\{#folder}\MakeManager.jar; WorkingDir: {%HOMEPATH}\{#folder}; Parameters:"t t"; Description: Iniciar Aplica��o;Flags: postinstall shellexec skipifsilent

[Code]
procedure FileReplaceString(const FileName, Searchfolder, Replacefolder, Searchjar, Replacejar: string);
var                                                        
  MyFile : TStrings;
  Myfolder, Myjar : string;
begin
  MyFile := TStringList.Create;
  try
    try
      MyFile.LoadFromFile(FileName);
      Myfolder := MyFile.Text;      
      { Only save if text has been changed. }
      if StringChangeEx(Myfolder, Searchfolder, Replacefolder, True) > 0 then
      begin;
        MyFile.Text := Myfolder;
        MyFile.SaveToFile(FileName);
      end;  
      if StringChangeEx(Myfolder, Searchjar, Replacejar, True) > 0 then
      begin;
        MyFile.Text := Myfolder;
        MyFile.SaveToFile(FileName);
      end;

    except
     
    end;
  finally
    MyFile.Free;
  end;
end;