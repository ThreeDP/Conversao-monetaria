# API Rest: Conversão de moeda
> Essa api tem o objetivo de realizar converções monetárias baseada em parametros informados em sua chamada.

**Request:**
> Para realizar chamadas á api o seguinte formato é utilizado.

```
    http://localhost:8000/exchange/{amount}/{from}/{to}/{rate}
    http://localhost:8000/exchange/10/BRL/USD/4.50

```

**Response:**
> A resposta da requisição segue os seguintes padrões.

```
{
  "valorConvertido": 45,
  "simboloMoeda": "$"
}

* Conversões:
 * De Real para Dólar;
 * De Dólar para Real;
 * De Real para Euro;
 * De Euro para Real;
 * De BTC para Dolar;
 * De BTC para Real;

* Salvar os dados no banco de dados:
 * criar uma rotina para salvar o dados para consultas futuras
```

# BUILD: Como executar

**Construção do executavel:**
> Para a construção do executavel.

```
make 
```