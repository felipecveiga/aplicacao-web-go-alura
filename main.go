package main

import "net/http" // pacote que cria servidores http


//O que é o handler ? O Handler é a funcao que criaremos para executar determinada ação quando acessarmos o caminho http://localhost:8080.
//Como será responsavel por lidar com as requisições e respostas do servidor, a função handler deverá receber dois parametros, um do tipo w http.ResponseWriter e outro do tipo r *http.Request.
// O primeiro parametro é responsavel por escrever uma resposta de volta ao clinete que fez a solicitação http, usamos ele para enviar dados de resposta ao navegador ou cliente que fez a solicitação.
// O segundo parametro é responsavel por receber a solicitação e fazer o tratamento da mesma, é como se fosse uma caixa que contem um monte de informações sobre uma solicitação que chega ao servidor.
func handler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w,r, "./static/index.html")// O http.ServeFile é uma funcao que facilita a entrega de arquivos estaticos em um servidor web, ele ajuda enviar esses arquivos para o navegador do usuario quando ele faz a requsicao da pagina ou recurso do seu servidor, no nosso caso enviar o arquivo index.html.
											//Ou seja, quando o usuario acessa a raiz do projeto "/" que é o localhost:8080, a funcao http.ServerFile é chamada para enviar o arquivo index.html para o navegador do usuario.
}

func main() {

	//A funcao HandleFunc recebe 2 parametros, o primeiro informa que a url http://localhost:8080 deverá executar determinada ação.
	//O segundo parametro indica que chamaremos a funcao handler ao acessar o endereco http://localhost:8080
	http.HandleFunc("/", handler)



	//o http.ListenAndServe é um comando que inicia um servidor web na maquina local, ele recebe pedidos, solicitacoes http e envia as resposta desse pedido.
	// O 8080 especifica a porta que o servidor vai escutar os pedidos.
	//O nil informa que utilizaremos a configuracao padrao do go.
	http.ListenAndServe(":8080", nil)// ele é responsavel por especificar em qual porta nossa aplicação irá rodar

}