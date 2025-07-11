# sensor-data-sim

Simulador de sensores IoT que gera dados realísticos e os envia via gRPC para um servidor de coleta de dados. Projetado para ser executado em Kubernetes como múltiplos pods simulando uma rede de sensores distribuída.

## 🚀 Características

- **Simulação Realística**: Gera dados de temperatura, pressão, umidade, ruído, luminosidade, qualidade do ar, vibração e métricas de bateria
- **Variações Temporais**: Simula ciclos diários, sazonais e variações ambientais naturais
- **Comunicação gRPC**: Protocolo eficiente para comunicação com o servidor
- **Containerizado**: Pronto para deploy em Kubernetes
- **Configurável**: Parâmetros ajustáveis via variáveis de ambiente
- **Métricas Customizadas**: Suporte a métricas adicionais específicas do domínio

## 📊 Dados Simulados

### Sensores Ambientais
- **Temperatura**: Variação diária e sazonal com ruído realístico
- **Pressão Atmosférica**: Variações baseadas em condições meteorológicas
- **Umidade**: Correlacionada inversamente com temperatura
- **Qualidade do Ar**: Índice AQI com variações por horário

### Sensores de Ruído e Luz
- **Nível de Ruído**: Variações baseadas em horário (mais ruído durante o dia)
- **Intensidade Luminosa**: Simulação do ciclo solar e luz artificial
- **Índice UV**: Radiação ultravioleta
- **Vibração**: Correlacionada com níveis de ruído

### Sistema de Energia
- **Tensão da Bateria**: 3.0V - 4.2V
- **Percentual de Carga**: 20% - 100%
- **Status de Carregamento**: Simulação realística
- **Temperatura da Bateria**: Correlacionada com temperatura ambiente

### Métricas Customizadas
- **CO2**: Concentração em PPM
- **Partículas de Poeira**: Densidade no ar
- **Velocidade do Vento**: m/s
- **Precipitação**: mm/h

## 🏗️ Arquitetura

```
┌─────────────────┐    gRPC     ┌─────────────────┐
│   Sensor Pod 1  │ ──────────► │                 │
├─────────────────┤             │                 │
│   Sensor Pod 2  │ ──────────► │  Servidor gRPC  │
├─────────────────┤             │                 │
│   Sensor Pod N  │ ──────────► │                 │
└─────────────────┘             └─────────────────┘
```

## 🛠️ Configuração

### Variáveis de Ambiente

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `SERVER_ADDR` | Endereço do servidor gRPC | `localhost:50051` |
| `SENSOR_ID` | ID único do sensor | `sensor-{timestamp}` |
| `SEND_INTERVAL` | Intervalo de envio | `5s` |
| `SENSOR_LAT` | Latitude | `-23.5505` |
| `SENSOR_LNG` | Longitude | `-46.6333` |
| `SENSOR_ALT` | Altitude (metros) | `760` |
| `SENSOR_LOCATION` | Descrição da localização | `São Paulo, Brasil` |

## 🚀 Como Usar

### Pré-requisitos

- Go 1.21+
- Docker (opcional)
- Kubernetes (opcional)
- Protocol Buffers compiler (para gerar código)

### Instalação

```bash
# Clonar repositório
git clone <seu-repositorio>
cd sensor-data-sim

# Instalar dependências
make deps

# Instalar ferramentas de desenvolvimento
make install-tools

# Gerar código protobuf
make proto
```

### Execução Local

```bash
# Executar um sensor
make run

# Executar múltiplos sensores para teste
make run-multiple

# Compilar e executar
make build
./bin/sensor-sim
```

### Execução com Docker

```bash
# Construir imagem
make docker-build

# Executar container
make docker-run
```

### Deploy no Kubernetes

```bash
# Deploy
make k8s-deploy

# Verificar status
make k8s-status

# Monitorar logs
make k8s-logs

# Remover
make k8s-undeploy
```

## 📁 Estrutura do Projeto

```
sensor-data-sim/
├── cmd/
│   └── sensor/
│       └── main.go              # Aplicação principal
├── pkg/
│   └── sensor/
│       └── simulator.go         # Lógica de simulação
├── proto/
│   └── sensor.proto             # Definições protobuf
├── k8s-deployment.yaml          # Configuração Kubernetes
├── Dockerfile                   # Imagem Docker
├── Makefile                     # Comandos de automação
├── go.mod                       # Dependências Go
└── README.md                    # Documentação
```

## 🔧 Desenvolvimento

### Comandos Úteis

```bash
# Formatar código
make fmt-fix

# Executar testes
make test

# Verificar código
make lint

# Limpar arquivos gerados
make clean
```

### Adicionando Novos Sensores

1. Edite `proto/sensor.proto` para adicionar novos campos
2. Regenere o código: `make proto`
3. Atualize `pkg/sensor/simulator.go` com a nova lógica
4. Teste as mudanças: `make test`

## 🌟 Exemplos de Uso

### Sensor Individual
```bash
export SENSOR_ID="sensor-lab-01"
export SENSOR_LOCATION="Laboratório Principal"
export SEND_INTERVAL="30s"
go run ./cmd/sensor
```

### Múltiplos Sensores
```bash
# Terminal 1 - Sensor externo
SENSOR_ID=outdoor-01 SENSOR_LAT=-23.5505 SENSOR_LNG=-46.6333 go run ./cmd/sensor &

# Terminal 2 - Sensor interno
SENSOR_ID=indoor-01 SENSOR_LAT=-23.5515 SENSOR_LNG=-46.6343 go run ./cmd/sensor &
```

## 🐳 Docker

### Construir Imagem
```bash
docker build -t sensor-data-sim:latest .
```

### Executar Container
```bash
docker run -e SERVER_ADDR=servidor:50051 \
           -e SENSOR_ID=sensor-docker-01 \
           -e SEND_INTERVAL=10s \
           sensor-data-sim:latest
```

## ☸️ Kubernetes

O projeto inclui configurações para deploy em Kubernetes com múltiplos pods simulando uma rede de sensores distribuída.

### Escalar Sensores
```bash
kubectl scale deployment sensor-simulator --replicas=10
```

### Monitorar
```bash
kubectl get pods -l app=sensor-simulator
kubectl logs -f deployment/sensor-simulator
```

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

## 🔮 Próximos Passos

- [ ] Implementação do servidor gRPC receptor
- [ ] Dashboard web para visualização dos dados
- [ ] Persistência em banco de dados
- [ ] Alertas baseados em thresholds
- [ ] Métricas de observabilidade (Prometheus)
- [ ] Simulação de falhas de rede
- [ ] Configuração via API REST
