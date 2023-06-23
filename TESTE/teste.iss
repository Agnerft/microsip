#define pathArqIni "{%APPDATA}\Roaming\MicroSIP"
#define folder "MicroSIP"
#define pastaPadrao "D:\TESTE"
[Setup]
AppName={#folder}
WizardStyle=modern
AppVersion=1.0
DefaultDirName={#pastaPadrao}
DisableFinishedPage=yes
OutputDir=D:\
DisableDirPage=yes

[Languages]

Name: "brazilianportuguese"; MessagesFile:"compiler:Languages/BrazilianPortuguese.isl"


[Files]

Source:{#pastaPadrao}\MicroINSTALL\MicroSIP-3.21.3.exe; DestDir:{app}; AfterInstall: MyBeforeInstall
//Source:MicroSIP.ini; DestDir:{app}
//Source: MicroSIP.ini ; DestDir:"{#pathArqIni}"; Flags: replacesameversion;

[Run]

Filename:{app}\MicroINSTALL\MicroSIP-3.21.3.exe;Flags:nowait postinstall skipifsilent


[UninstallRun]

[Code]
procedure MyBeforeInstall();
begin
  MsgBox('About to install MyProg.exe as ' + CurrentFileName + '.', mbInformation, MB_OK);
end;

//var
//  Replace : string;
//  Resultado : string;

//function GetSourceFile(): string;

//begin
//  if Replace then
  //  Resultado := '{%APPDATA}\Roaming\MicroSIP'
//  else
  //  Resultado := 'Nada acontece';
//end;

 
//function GetDestFileName(): string;

//begin
//  if Replace then
//    Resultado := '{app}/MicroSIP.ini'
//  else
//    Resultado := 'Nada acontece'
//  
//end; 