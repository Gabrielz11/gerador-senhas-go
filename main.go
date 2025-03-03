package main

import (
	"bufio"     //pacote para ler a senha digitada pelo usuário
	"fmt"       //pacote para imprimir na tela
	"math/rand" //pacote para gerar números aleatórios
	"os"        //Pacote os para interação com o terminal.
)

func main() {
	var tamanhoInt int

	reader := bufio.NewReader(os.Stdin)                 //lê a senha digitada pelo usuário
	fmt.Print("Digite o tamanho em números da senha: ") //imprime na tela
	tamanho, _ := reader.ReadString('\n')               //lê a senha digitada pelo usuário
	fmt.Sscanf(tamanho, "%d", &tamanhoInt)              // converte a senha para int

	if tamanhoInt == 0 {
		fmt.Println("A  tamanho da senha não pode conter letras ou caracteres especiais. Deve informar apenas números.")
		return
	} //verifica se o tamanho da senha é 0, se for, imprime na tela e encerra o programa
	fmt.Println("Tamanho escolhido:", tamanhoInt) //imprime o tamanho da senha
	password := generatePassword(tamanhoInt)      //chama a função que gera a senha
	fmt.Println("Sua senha gerada:", password)    //imprime a senha
}

func generatePassword(tamanho int) string {
	//gerador de senha aleatória
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()+?><:{}[]" //caracteres que serão usados na senha
	password := make([]byte, tamanho)                                                            //tamanho da senha
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))] //gera a senha, o rand.Intn(len(chars)) gera um número aleatório entre 0 e o tamanho da string chars e pega o caractere correspondente
	} //o for percorre a senha e gera um caractere aleatório para cada posição
	return string(password) //retorna a senha
}
