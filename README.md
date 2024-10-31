# Implementação de Criptografia Assimétrica com RSA em Go

## Introdução

Esta implementação demonstra como cifrar e decifrar mensagens usando criptografia assimétrica com o algoritmo RSA. O RSA utiliza um par de chaves: uma chave pública para cifrar a mensagem e uma chave privada para decifrá-la. 

## Funcionamento

1. **Geração de Chaves**: Um par de chaves RSA é gerado, composto por uma chave pública e uma chave privada.
2. **Cifragem**: A mensagem é cifrada utilizando a chave pública.
3. **Decifragem**: A mensagem cifrada é decifrada utilizando a chave privada.
4. **Exportação das Chaves**: As chaves são exportadas em formato PEM para uso futuro.

## Dependências

- Go 1.15 ou superior

## Execução

Para executar o código, utilize o seguinte comando no terminal:

```bash
go run main.go
```
![Screenshot](screenshot.png)


