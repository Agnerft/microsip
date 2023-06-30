#define pathArqIni "{%APPDATA}\Roaming\MicroSIP"
#define folder "MicroSIP"
#define pastaPadrao "D:\TESTE"
#define outputDir "D:\CRIA"
[Setup]
AppName={#folder}
WizardStyle=modern
AppVersion=1.0
DefaultDirName={#pastaPadrao}
DisableFinishedPage=yes
OutputDir={#outputDir}
DisableDirPage=yes

[Languages]

Name: "brazilianportuguese"; MessagesFile:"compiler:Languages/BrazilianPortuguese.isl"


[Files]

Source:{#pastaPadrao}\MicroINSTALL\MicroSIP-3.21.3.exe; DestDir:{#outputDir}\MicroINSTALL; Flags: uninsneveruninstall onlyifdoesntexist ignoreversion
Source:{#pastaPadrao}\MicroINI\MicroSIP.ini; DestDir:{#outputDir}\MicroINI; Flags: uninsneveruninstall onlyifdoesntexist ignoreversion
//Source:MicroSIP.ini; DestDir:{app}
//Source: MicroSIP.ini ; DestDir:"{#pathArqIni}"; Flags: replacesameversion;

[Run]

Filename:{app}\MicroINSTALL\MicroSIP-3.21.3.exe;Flags:nowait postinstall skipifsilent



[Code]

[Code]
[Code]
var
  InputPage: TInputQueryWizardPage;
  InputValue1: string;
  InputValue2: string;

procedure InitializeWizard;
begin
  // Cria uma página de consulta ao usuário
  InputPage := CreateInputQueryPage(wpSelectDir, 'Preenchimento', 'Digite um valor:', 'Insira um valor para preenchimento.');

  // Adiciona o campo de preenchimento
  InputPage.Add('Sip:', False);
  InputPage.Add('link da Telefonia', False);
end;

function NextButtonClick(CurPageID: Integer): Boolean;
begin
  if CurPageID = InputPage.ID then
  begin
    // Ação a ser executada ao clicar em Avançar na página de consulta ao usuário
    InputValue1 := InputPage.Values[0];
    InputValue2 := inputPage.Values[1]; 
    MsgBox('Valor do Sip é: ' + InputValue1 + 'E o link da GVC é: ' + InputValue2, mbInformation, MB_OK);
    
  end;

  Result := True;
end;


