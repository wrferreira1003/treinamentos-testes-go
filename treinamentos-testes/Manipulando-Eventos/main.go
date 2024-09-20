package main

func main() {

}

/*
- Evento é algo que aconteceu no passado e que pode ser processado no futuro.
- Eventos são importantes para o desenvolvimento de aplicações reais, pois nos permitem
  desacoplar o código, ou seja, podemos processar um evento em um momento diferente do
  momento em que ele ocorreu.

Ex: Inseri um novo cliente no meu sistema. agora o que quero fazer apos isso?
  - Enviar um email para o cliente com um link para o sistema.
  - Publicar uma mensagem na fila
  - Notificar um usuario.

  A ideia quando acontece um evento podemos disparar uma funcao que vai processar esse evento.

  Ex:
    - Inserir cliente
    - Enviar email

  Quando acontece um evento, uma funcao e disparada para processar esse evento.

  Elemento Tatico de um contexto de eventos:
	- elemento principal: Seria o evento.
	- Operacoes que serao executadas quando o evento ocorrer.
	- Gerenciado dos eventos/operacoes, responsavel por garantir que as operacoes sejam executadas de forma correta.
	- Disparo de eventos.

*/
