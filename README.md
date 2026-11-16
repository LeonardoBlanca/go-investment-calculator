# Investment Calculator

Calculadora de investimentos de linha de comando escrita em Go. Estima quanto um
investimento inicial vale após um número de anos, usando juros compostos, e
mostra também esse valor descontada a inflação.

## Requisitos

- Go 1.24 ou superior

## Como executar

```bash
go run .
```

Ou gere um executável:

```bash
go build -o investment-calculator .
./investment-calculator
```

O programa pede três valores:

| Pergunta                   | Significado                              | Exemplo |
|----------------------------|------------------------------------------|---------|
| `Investment Amount`        | Valor investido inicialmente             | `1000`  |
| `Expected Return Rate (%)` | Taxa de retorno anual esperada, em %     | `5.5`   |
| `Years`                    | Por quantos anos o dinheiro fica aplicado | `10`    |

Use ponto como separador decimal (`5.5`, não `5,5`).

## Exemplo

```
Investment Amount: 1000
Expected Return Rate (%): 5.5
Years: 10
Future Value: 1708.14
Future Value (adjusted for inflation): 1334.40
```

## Como o cálculo funciona

**Valor futuro** (juros compostos):

```
VF = P × (1 + r/100)^n
```

**Valor futuro real** (descontada a inflação):

```
VR = VF / (1 + i/100)^n
```

Onde:

- `P` é o valor investido
- `r` é a taxa de retorno anual (%)
- `n` é o número de anos
- `i` é a taxa de inflação anual (%), fixa em **2,5%** no código (`inflationRate`)

## Limitações

- A inflação é fixa; para alterá-la, edite a constante `inflationRate` em
  `investment_calculator.go`.
- Considera apenas um aporte inicial, sem aportes mensais.
- Não há validação de entrada: valores não numéricos ou negativos não geram aviso.
- Não desconta impostos nem taxas.
