# Definizione struttura progetto

## Nuova struttura package

```
├── cmd
|   └── main.go                  main principale
|   └── main-test.go             test relativi al main principale
│   └── server                   contiene le api (ex controller)
│        └── swagger-api.go      contiene le api per esporre la documentazione swagger
│        └── template-api.go     contiene le api per esporre la risorsa REST
│        └── middleware          contiene gli handler comuni (log, jwt, health)
|            └── audit           modulo per la gestione audit (attivita utente)
|            └── authentication  modulo per la gestione dell'autenticazione (JWT, ruoli, etc)
├── pkg
│    ├── broker                  modulo per l'utilizzo di kafka (producer e consumer)
│    ├── config                  modulo per l'accesso alla configurazione remota
|    ├── constants               costanti da utilizzare nel progetto
|    ├── db                      modulo per l'implementazione dell'accesso al db (mongo)
|    ├── mutex                   modulo per la gestione dei semafori (lock e unlock) su db (mongo)  
│    ├── template                modulo che implementa la risorsa rest (template)
│    ├── utils                   utils varie (serializzazione, json, calcolo hash e http params)
│    ├── errors                  modulo per la gestione degli errori
```

## How To Use

1. Copiare il template di progetto al path desiderato es /home/go/src/gitlab.tdnet.it/<gruppo>/mio-progetto-ms

2. All'interno della cartella lanciare ./istanziaTemplate.sh per la generazione del progetto e seguire la procedura di generazione   progetto

3. Caricare la configurazione su config server con il comando:
   curl -X POST --header 'Content-Type: multipart/form-data' --header 'Accept: application/json' --header 'X_RTSYSID: XXX_001' -F "file=@config.toml" 'http://10.1.3.106:7892/configuration-extra/toml?name=mio-progetto-ms&profile=svil'
4. Lanciare lo script build.sh per la build del progetto
   In alternativa si può compilare mediante il seguente comando:
      go build -o ```<nome eseguibile>``` cmd/main.go
5. Per scaricare le librerie in locale eseguire il seguente comando:
    go mod vendor
6. Per lanciare il progetto eseguire il seguente comando:
    ```./<nome-eseguibile>```  o in alternativa
    ```go run cmd/main.go --profile <nome-profilo> dove <nome-profilo>``` può assumere i seguenti valori:
      - local
      - svil
      - stage
      - prod
