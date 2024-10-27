# UOW

- Unit of Work
- Pattern de projeto
- Gerenciamento de transações

## Exemplo de uso
    Estamos com um use-case onde precisa ser feito duas persistencia de dados, uma em um banco de dados e outra em outro.
    Precisamos que ambas as persistências ocorram dentro da mesma transação, ou seja, ambas devem ter sucesso ou nenhuma deve ser persistida.

    Para isso vamos usar o padrão de projeto UOW.

    - Precisamos fazer:
        - BeginTransaction
        - Commit
        - Rollback

    Como estamos trabalhando com repositorios diferentes, como conseguir gerenciar a transação?

    É responsabilidade do use-case de gerenciar a transação?
      - Sim, o use-case é responsável por gerenciar a transação.
  
    Entendendo Unit of Work
      - É um padrão de projeto que gerencia as transações.
      - Gerencia as transações de forma transacional, ou seja, todas as operações devem ter sucesso ou nenhuma deve ser persistida.
      - Gerencia as transações de forma explícita, ou seja, o desenvolvedor deve gerenciar as transações.
  
    BEGIN
        Transacao 1 -> Repo
        Transacao 2 -> Repo
    COMMIT
        Transacao 1 -> Repo
        Transacao 2 -> Repo
    ROLLBACK
        Transacao 1 -> Repo
        Transacao 2 -> Repo

    Nesse caso quem vai gerar os repositórios sera o Unit of Work.

    Precisamos registrar os repositórios que serão gerados pelo Unit of Work.
      Register
      GetRepository
      Unregister
      DO(
        Transacao 1 -> Repo
        Transacao 2 -> Repo 
        COMMIT - ROLLBACK
      )
