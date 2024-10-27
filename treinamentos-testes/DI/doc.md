# DI - Dependency Injection

## Conceitos

- **Inversão de Controle (IoC)**: É um padrão de projeto que permite que uma classe delegue a responsabilidade de criar e gerenciar objetos a outros objetos.
- **Injeção de Dependência (DI)**: É um padrão de projeto que permite que uma classe delegue a responsabilidade de criar e gerenciar objetos a outros objetos.

## Vantagens

- **Desacoplamento**: O código fica mais fácil de entender e manter, pois as dependências são gerenciadas por um container externo.
- **Facilidade na criação de mocks**: Facilita a criação de mocks para testes.
- **Testabilidade**: Facilita a criação de testes unitários.
- **Flexibilidade**: Permite a substituição de dependências em tempo de execução.

## Desvantagens

- **Complexidade**: Pode ser mais complexo de entender e manter, pois envolve a criação de um container externo.
- **Performance**: Pode ter um impacto na performance, pois envolve a criação de um container externo.

## Precisamos ter de alguma forma de resolver o problema de dependencias

- Projeto ficando grande, fica insustentavel manter as depedencias, e com isso precisamos usar o DI,
  que seria a inversão de controle, onde o container gerencia as dependencias.

### Bibliotecas de DI

- google wire - https://github.com/google/wire
