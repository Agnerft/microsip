#!/bin/bash

ls -d */ > files_tmp.txt



arq_txt=files_tmp.txt
subpastas=("CRM_ESP_1" "CRM_OI" "CRM_CLARO" "CRM_NET" "CRM_SKY" "CRM_VIVO" "CRM_BASIC")

# Loop através das pastas listadas no arquivo
while IFS= read -r pasta; do
    echo "Entrando em $pasta"
    cd "$pasta"  # Muda para a pasta, sai do script se falhar

    # Loop através das subpastas específicas
    for subpasta in "${subpastas[@]}"; do
        echo "Entrando em $subpasta"
        cd "$subpasta" # Muda para a subpasta, sai do script se falhar

        echo "Entrando em mailings_uploads"
        cd mailings_uploads        
        # Remove arquivos que contêm "_INVALIDOS.csv" no final do nome
        find . -type f -name "*_INVALIDOS.csv" -exec rm -f {} \;
        
        cd ..
        cd ..                # Volta para a subpasta anterior
    done
    
    cd ..                      # Volta para a pasta principal
done < "$arq_txt"


rm /opt/payara/glassfish/domains/domain1/docroot/files/$arq_txt
