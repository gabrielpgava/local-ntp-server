# NTP Server em Go

Este projeto implementa um servidor NTP (Network Time Protocol) simples utilizando a linguagem Go. Ele permite que dispositivos em uma rede local consultem a data e hora do servidor, ideal para ambientes de laboratório, testes ou redes isoladas.

## Recursos

- **Servidor NTP básico:** Atende requisições NTP padrão (porta UDP 123) e responde com a hora local do sistema.
- **Log de acessos:** Registra cada consulta feita ao servidor, incluindo data e hora, IP do cliente e nome do host (quando possível), em um arquivo `ntp_access.log`.
- **Código simples e didático:** Fácil de entender, modificar e expandir para necessidades locais ou laboratoriais.

## Avisos

> **Importante:** Este servidor NTP é destinado apenas para uso em ambientes controlados, redes locais ou para fins educacionais. Não utiliza autenticação, não sincroniza com fontes externas de tempo e não implementa todas as funcionalidades do NTP padrão.  
> **Não utilize em ambientes de produção ou onde a precisão e segurança do tempo sejam críticas.**

## Como Buildar

O projeto é escrito em Go puro e não possui dependências externas além da biblioteca padrão.

### Pré-requisitos

- Go 1.16 ou superior instalado ([download aqui](https://go.dev/dl/))

### Build para cada plataforma

#### Linux

```sh
go build -o ntp-server main.go
```

#### Windows

```sh
go build -o ntp-server.exe main.go
```

#### macOS

```sh
go build -o ntp-server main.go
```

#### Cross-compile para outras plataformas

Você pode compilar para diferentes sistemas operacionais e arquiteturas utilizando as variáveis de ambiente do Go:

- **Para Windows 64 bits:**
  ```sh
  GOOS=windows GOARCH=amd64 go build -o ntp-server.exe main.go
  ```
- **Para Linux 64 bits:**
  ```sh
  GOOS=linux GOARCH=amd64 go build -o ntp-server main.go
  ```
- **Para macOS ARM64 (Apple Silicon):**
  ```sh
  GOOS=darwin GOARCH=arm64 go build -o ntp-server main.go
  ```

Veja todas as opções de sistemas e arquiteturas suportadas em: https://golang.org/doc/install/source#environment

## Como Usar

1. Compile conforme acima para sua plataforma.
2. Execute como administrador/root (porta 123 é privilegiada):
   ```sh
   sudo ./ntp-server
   ```
   No Windows, execute o terminal como administrador.

3. Clientes NTP podem agora consultar a hora do servidor usando o IP/máquina onde o `ntp-server` está rodando.

4. Consulte o arquivo `ntp_access.log` para visualizar os acessos realizados.

## Exemplo de Log

```
2025/05/24 16:01:12 Hora consultada em 2025-05-24 16:01:12 por 192.168.0.10 (hostname: workstation-10.local)
2025/05/24 16:05:43 Hora consultada em 2025-05-24 16:05:43 por 192.168.0.23 (hostname: -)
```

---

Feito com ❤️ pela Bi-Ga Tech