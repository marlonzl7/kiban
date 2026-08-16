# kiban

CLI de setup de ambiente de desenvolvimento para Linux. Declarativa,
orientada a arquivos YAML, cross-distro.

Defina as ferramentas que você usa uma vez, num arquivo `setup.yaml`, e
instale tudo com um único comando, em qualquer máquina, em qualquer
distro suportada.

```bash
kiban install --file setup.yaml
```

## Por quê

Formatar a máquina, trocar de computador, entrar em um projeto novo:
o processo de configurar o ambiente de desenvolvimento é sempre o mesmo,
manual e repetitivo. O kiban resolve isso com um arquivo declarativo que
você mantém no seu GitHub e roda sempre que precisar.

## Instalação

Baixe o binário mais recente na [página de releases](https://github.com/marlonzl7/kiban/releases).

```bash
curl -L -o kiban https://github.com/marlonzl7/kiban/releases/latest/download/kiban_v0.2.0_linux_amd64
chmod +x kiban
sudo mv kiban /usr/local/bin/
```

Confirme a instalação:
```bash
kiban --version
```

## Uso

Crie um `setup.yaml` na raiz do seu projeto (ou mantenha um pessoal em
qualquer lugar):

```yaml
version: 1
tools:
  containers:
    - docker
  utils:
    - git
    - curl
    - zsh
  languages:
    - java@21
```

Instale:
```bash
kiban install --file setup.yaml
# ou, se o arquivo se chama setup.yaml e está no diretório atual:
kiban install
```

O kiban detecta automaticamente sua distro, arquitetura e gerenciador de
pacotes, valida o arquivo, confirma a sessão de sudo uma única vez, e
instala cada ferramenta na ordem definida, reportando o progresso e um
resumo final:
```bash
Environment: ubuntu (apt) | x86_64 | sudo available
Valid setup. Starting tool installation...
Installing docker... [OK]
Installing git... [OK]
Installing curl... [OK]
Installing zsh... [OK]
Installing java... [OK]
Summary: 5 installed, 0 failed
```

### Versionamento de ferramentas

Ferramentas que suportam versão específica aceitam a sintaxe `ferramenta@versão`:

```yaml
languages:
  - java@21
```

Se a versão for omitida e a ferramenta tiver uma versão padrão definida
no catálogo, ela é usada automaticamente. Se não houver versão nem
padrão disponível, o `kiban` avisa antes de instalar qualquer coisa.

## Catálogo de ferramentas suportadas

O catálogo completo, com o schema de cada ferramenta, fica em
[`internal/loader/tools/`](internal/loader/tools/), organizado por categoria
(`containers/`, `utils/`, `languages/`, `editors/`, `databases/`). Cada
arquivo `.yaml` documenta os passos de instalação por distro/gerenciador
de pacotes.

Distros suportadas: Ubuntu, Fedora, Arch (e derivadas, via `ID_LIKE`).

## Limitações conhecidas

- **`java@<versão>` pode falhar no Fedora**: o Fedora remove versões
  antigas de OpenJDK dos repositórios padrão conforme lança novas (ex:
  `java-21-openjdk` deixou de existir a partir do Fedora 44). Se a
  instalação falhar com "nenhuma correspondência", ajuste a versão
  solicitada no `setup.yaml` para uma disponível na sua versão do
  Fedora, ou omita a versão (`java` em vez de `java@21`) e ajuste
  `default_version` em `tools/languages/java.yaml`.

- **Node no Arch (pacman) ignora a versão solicitada**: o pacote
  `nodejs` não é versionado por major nesse PM, diferente de apt/dnf —
  sempre instala a versão mais recente do repositório oficial,
  independente de `node@<versão>` no `setup.yaml`.

- **Redis pode resolver para Valkey**: Fedora e Arch já substituem o
  pacote `redis` por Valkey (fork open-source compatível) nos
  repositórios padrão; Ubuntu ainda mantém Redis genuíno. A
  verificação pós-instalação aceita ambos.

- **GPG pode falhar silenciosamente na primeira execução em VM
  totalmente nova** (instalação do Node via apt): o keyring pode ficar
  vazio/não convertido na primeira chamada de `gpg --dearmor`,
  causando erro de assinatura no `apt-get update` seguinte.
  Reexecutar `kiban install` resolve.

## Aviso de segurança

O kiban executa comandos de instalação (incluindo `sudo`) definidos nos
arquivos YAML do catálogo. O kiban **não audita nem hospeda** o conteúdo
instalado por cada ferramenta, a responsabilidade sobre a segurança e
integridade do software instalado é do fornecedor de cada pacote/repositório
upstream (ex: o pacote oficial do Docker, do OpenJDK, etc).

Antes de instalar, recomendo revisar o arquivo YAML da ferramenta no
catálogo (`internal/loader/tools/`) para confirmar exatamente quais
comandos serão executados no seu sistema.

## Contribuindo

Quer adicionar uma nova ferramenta ou distro? Veja o [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Status

Projeto em desenvolvimento ativo (MVP). Escopo atual: Linux only
(Ubuntu, Fedora, Arch). macOS e Windows estão no roadmap.
