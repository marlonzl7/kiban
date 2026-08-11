# Contribuindo com o kiban
 
Obrigado por considerar contribuir com o kiban! A forma mais comum de
contribuição é adicionar suporte a uma ferramenta nova ou a uma distro
que ainda não é suportada — e para isso, **você não precisa saber Go**.
 
Toda ferramenta do catálogo é definida em um arquivo YAML, que descreve
como instalar e verificar aquela ferramenta. Editar ou criar um desses
arquivos é o suficiente para contribuir.
 
Este guia explica o passo a passo para adicionar uma ferramenta nova, o
schema que todo arquivo YAML deve seguir, e as regras de aceitação que
usamos para manter o catálogo consistente.
 
## Como adicionar uma ferramenta nova
 
1. **Escolha a categoria e crie o arquivo**
   Veja a [estrutura de pastas](#estrutura-de-pastas) e decida em qual
   categoria a ferramenta se encaixa (`languages/`, `containers/`,
   `editors/`, `utils/`, `databases/`). Crie o arquivo em
   `tools/<categoria>/<ferramenta>.yaml`.
2. **Copie a estrutura base**
   Use um arquivo existente do catálogo como modelo (ex:
   `tools/utils/git.yaml`) para garantir que todas as chaves
   obrigatórias do schema estejam presentes.
3. **Preencha nome e descrição**
   `name` é o nome do pacote (veja a Regra 3 abaixo — pode divergir do
   nome do binário). A descrição deve ser curta e direta.
4. **Defina os steps de instalação**
   Escreva os comandos de instalação para cada package manager
   suportado (`apt`, `dnf`, `pacman`). Sempre que possível, teste os
   comandos em uma VM real antes de abrir o PR — principalmente se você
   nunca instalou a ferramenta daquela forma antes.
5. **Defina a verificação (`verify`)**
   Adicione um comando que confirme que a instalação funcionou (ex:
   `<ferramenta> --version`) e o texto esperado na saída.
6. **Confira as regras de aceitação**
   Antes de abrir o PR, revise a seção [Regras de aceitação](#regras-de-aceitação)
   abaixo — elas cobrem casos como métodos de instalação concorrentes,
   comandos não-interativos e meta-pacotes.
## Schema do arquivo YAML
 
Todo arquivo de ferramenta segue esta estrutura:
 
| Campo             | Obrigatório | Descrição |
|--------------------|-------------|-----------|
| `name`             | Sim         | Nome do **pacote** a ser instalado (não necessariamente o nome do binário — veja Regra 3) |
| `description`      | Sim         | Descrição curta do que a ferramenta faz |
| `category`         | Sim         | Categoria da ferramenta (deve corresponder à pasta onde o arquivo está) |
| `version_flag`     | Sim         | Comando usado para checar a versão instalada (ex: `git --version`) |
| `default_version`  | Não         | Versão usada quando o usuário não especifica uma no `setup.yaml`. Ausente/vazio = a ferramenta sempre instala a versão mais recente do repositório |
| `install`          | Sim         | Um bloco por package manager suportado (`apt`, `dnf`, `pacman`), cada um com uma lista de `steps` |
| `install.<pm>.steps[].cmd`  | Sim | Comando de instalação. Deve ser não-interativo (veja Regra 2) |
| `install.<pm>.steps[].sudo` | Sim | `true` se o comando precisa rodar com sudo |
| `verify.cmd`       | Sim         | Comando executado após a instalação para confirmar que funcionou |
| `verify.expect`    | Sim         | Texto que deve aparecer na saída de `verify.cmd` para considerar sucesso |
| `post_install.message` | Sim (pode ser vazio: `""`) | Mensagem exibida ao usuário após a instalação (ex: instruções manuais adicionais) |
 
### Exemplo comentado
 
```yaml
name: java                     # nome do pacote, usado no setup.yaml (ex: java@21)
description: Java Development Kit
category: languages            # deve bater com a pasta: tools/languages/
version_flag: java -version
default_version: "21"          # usado se o usuário não especificar versão
 
install:
  apt:                         # um bloco por package manager suportado
    steps:
      - cmd: apt-get install -y openjdk-{{version}}-jdk   # {{version}} é resolvido em tempo de execução
        sudo: true              # sempre true quando o comando precisa de privilégios elevados
 
verify:
  cmd: java -version
  expect: "openjdk version"    # texto esperado na saída do verify.cmd
 
post_install:
  message: ""                  # pode ficar vazio se não houver instrução adicional
```
 
## Estrutura de pastas
 
Os arquivos de ferramentas ficam em `internal/loader/tools/`, organizados
por categoria:
 
```
internal/loader/tools/
├── languages/     # java, node, python, go, rust...
├── containers/    # docker, podman...
├── editors/       # neovim, vscode...
├── utils/         # git, curl, zsh, ripgrep...
└── databases/     # postgresql, mysql, redis...
```
 
O caminho do arquivo deve seguir o padrão
`tools/<categoria>/<ferramenta>.yaml`, e o campo `category` dentro do
YAML deve corresponder ao nome da pasta (ex: um arquivo em
`tools/utils/htop.yaml` deve ter `category: utils`).
 
Se a ferramenta não se encaixa em nenhuma categoria existente, abra uma
discussão no PR antes de criar uma pasta nova.
 
## Regras de aceitação
 
Estas regras existem para manter o catálogo consistente e confiável.
PRs que não seguirem essas regras podem ser solicitados a ajustes antes
do merge.
 
### Regra 1 — Precedência de métodos de instalação
 
Quando existe mais de um jeito de instalar a mesma ferramenta, siga esta
ordem de preferência:
 
1. **Repositório oficial do mantenedor**, gerenciado via `apt`/`dnf`/`pacman`
   (ex: repositório oficial do Docker)
2. **Repositório da distro** (ex: `apt-get install docker.io`)
3. **Script de terceiros/shell script oficial** (ex: `curl | sh`)
O critério de desempate entre os níveis 1 e 2 é a atualidade e
confiabilidade da fonte — ambos garantem idempotência pelo próprio
package manager. O nível 3 é último recurso e exige que o contribuidor
verifique manualmente que o script é seguro para reexecução, já que não
há garantia estrutural de idempotência nesse caso.
 
### Regra 2 — Comandos não-interativos
 
Todo `cmd` em um arquivo de ferramenta precisa rodar sem exigir
confirmação do usuário. Use sempre as flags de instalação silenciosa:
 
```yaml
cmd: apt-get install -y git      # correto
cmd: apt-get install git         # incorreto — pede confirmação e trava a execução
```
 
O kiban não conecta entrada padrão ao processo executado, então comandos
interativos travam ou abortam silenciosamente durante a instalação.
 
### Regra 3 — `name` é o nome do pacote, não do binário
 
O campo `name` deve refletir o nome do **pacote** instalado pelo package
manager, mesmo quando ele diverge do nome do binário resultante:
 
```yaml
# tools/utils/ripgrep.yaml
name: ripgrep          # nome do pacote
...
verify:
  cmd: rg --version    # nome do binário é diferente
  expect: "ripgrep"
```
 
### Regra 4 — Verificação de meta-pacotes
 
Para meta-pacotes (ex: `build-essential`, grupos de desenvolvimento), o
campo `verify` deve checar apenas o binário mais essencial do grupo, não
paridade exata entre distros. A instalação do pacote já é uma transação
atômica do package manager — não é necessário validar cada componente
individualmente.
 
### Regra 5 — Evite YAML redundante
 
Se um meta-pacote já cobre a ferramenta que você quer adicionar, não crie
um arquivo individual redundante, a menos que exista um caso de uso real
para instalar aquela ferramenta isoladamente (fora do meta-pacote).
 
### Regra 6 — Resolução obrigatória de `{{version}}`
 
Toda ferramenta que usa `{{version}}` no `cmd` precisa garantir que essa
versão seja resolvida antes da execução — seja pelo usuário informando
`ferramenta@versão` no `setup.yaml`, seja pelo campo `default_version`
no arquivo da ferramenta. Sem nenhum dos dois, a instalação falha antes
de começar, com uma mensagem clara em vez de um erro tardio no comando.
 
## Exemplos comentados
 
O exemplo de ferramenta versionada (`java.yaml`, usando `{{version}}` e
`default_version`) já foi mostrado na seção [Schema do arquivo YAML](#schema-do-arquivo-yaml)
acima. Abaixo, um segundo exemplo cobrindo a Regra 3 — nome do pacote
diferente do nome do binário:
 
```yaml
# tools/utils/ripgrep.yaml
name: ripgrep                  # nome do pacote instalado pelo apt/dnf/pacman
description: Fast command-line search tool for finding text in files
category: utils
version_flag: rg --version     # o binário se chama "rg", não "ripgrep"
 
install:
  apt:
    steps:
      - cmd: apt-get install -y ripgrep
        sudo: true
 
verify:
  cmd: rg --version             # verify usa o nome do binário
  expect: "ripgrep"
 
post_install:
  message: ""
```
 
## Aviso de segurança
 
Ao contribuir com um arquivo de ferramenta, você está definindo comandos
que serão executados — muitas vezes com `sudo` — na máquina de quem usa
o kiban. Isso implica uma responsabilidade direta na hora de propor uma
fonte de instalação:
 
- Prefira sempre a fonte mais confiável disponível, seguindo a
  [Regra 1](#regra-1--precedência-de-métodos-de-instalação): repositório
  oficial do mantenedor antes de repositório da distro, antes de script
  de terceiros.
- Se o método de instalação envolver um script externo (`curl | sh` ou
  similar), verifique a origem antes de propor o PR e garanta que ela é
  oficial e mantida pelo próprio fornecedor da ferramenta.
- Lembre-se que o kiban **não audita nem hospeda** o conteúdo instalado
  por cada ferramenta — a responsabilidade sobre a segurança e
  integridade do que é instalado é do fornecedor upstream, mas a
  responsabilidade de apontar para essa fonte corretamente é do
  contribuidor.
Em caso de dúvida sobre a confiabilidade de uma fonte, mencione isso
explicitamente na descrição do PR para que possa ser discutido antes do
merge.