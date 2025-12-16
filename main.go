package main

import (
	"fmt"
)

// TreeNode - структура узла двоичного дерева
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree - структура двоичного дерева
type BinaryTree struct {
	Root *TreeNode
}

// Insert - функция вставки элемента в двоичное дерево
func (bt *BinaryTree) Insert(value int) {
	bt.Root = insertRecursive(bt.Root, value)
}

// Вспомогательная рекурсивная функция для вставки
func insertRecursive(node *TreeNode, value int) *TreeNode {
	// Если узел пустой, создаем новый узел
	if node == nil {
		return &TreeNode{Value: value}
	}

	// Рекурсивно вставляем в левое или правое поддерево в зависимости от значения
	if value < node.Value {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = insertRecursive(node.Right, value)
	}
	// Если значение уже существует, не добавляем его снова

	return node
}

// Search - функция поиска элемента в двоичном дереве
func (bt *BinaryTree) Search(value int) bool {
	return searchRecursive(bt.Root, value)
}

// Вспомогательная рекурсивная функция для поиска
func searchRecursive(node *TreeNode, value int) bool {
	// Если узел пустой, значит элемент не найден
	if node == nil {
		return false
	}

	// Если значение совпадает с текущим узлом, возвращаем true
	if node.Value == value {
		return true
	}

	// Рекурсивно ищем в левом или правом поддереве в зависимости от значения
	if value < node.Value {
		return searchRecursive(node.Left, value)
	} else {
		return searchRecursive(node.Right, value)
	}
}

// Delete - функция удаления элемента из двоичного дерева
func (bt *BinaryTree) Delete(value int) {
	bt.Root = deleteRecursive(bt.Root, value)
}

// Вспомогательная рекурсивная функция для удаления
func deleteRecursive(node *TreeNode, value int) *TreeNode {
	// Если узел пустой, возвращаем nil
	if node == nil {
		return nil
	}

	// Находим узел для удаления
	if value < node.Value {
		node.Left = deleteRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteRecursive(node.Right, value)
	} else {
		// Нашли узел для удаления
		
		// Случай 1: Узел без потомков (листовой узел)
		if node.Left == nil && node.Right == nil {
			return nil
		}
		
		// Случай 2: Узел с одним потомком
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		
		// Случай 3: Узел с двумя потомками
		// Находим минимальный элемент в правом поддереве
		minNode := findMin(node.Right)
		// Заменяем значение текущего узла значением минимального узла
		node.Value = minNode.Value
		// Удаляем минимальный узел из правого поддерева
		node.Right = deleteRecursive(node.Right, minNode.Value)
	}

	return node
}

// findMin - находит минимальный узел в дереве (крайний левый)
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// PrintInOrder - функция печати элементов дерева в порядке возрастания (in-order обход)
func (bt *BinaryTree) PrintInOrder() {
	fmt.Print("Элементы дерева (in-order): ")
	printInOrderRecursive(bt.Root)
	fmt.Println()
}

// Вспомогательная рекурсивная функция для in-order обхода
func printInOrderRecursive(node *TreeNode) {
	if node != nil {
		printInOrderRecursive(node.Left)
		fmt.Printf("%d ", node.Value)
		printInOrderRecursive(node.Right)
	}
}

// PrintPreOrder - функция печати элементов дерева в pre-order порядке
func (bt *BinaryTree) PrintPreOrder() {
	fmt.Print("Элементы дерева (pre-order): ")
	printPreOrderRecursive(bt.Root)
	fmt.Println()
}

// Вспомогательная рекурсивная функция для pre-order обхода
func printPreOrderRecursive(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		printPreOrderRecursive(node.Left)
		printPreOrderRecursive(node.Right)
	}
}

// PrintPostOrder - функция печати элементов дерева в post-order порядке
func (bt *BinaryTree) PrintPostOrder() {
	fmt.Print("Элементы дерева (post-order): ")
	printPostOrderRecursive(bt.Root)
	fmt.Println()
}

// Вспомогательная рекурсивная функция для post-order обхода
func printPostOrderRecursive(node *TreeNode) {
	if node != nil {
		printPostOrderRecursive(node.Left)
		printPostOrderRecursive(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}

func main() {
	// Создаем новое двоичное дерево
	tree := &BinaryTree{}

	// Добавляем элементы в дерево
	values := []int{50, 30, 70, 20, 40, 60, 80}
	fmt.Println("Добавляем элементы в дерево:")
	for _, v := range values {
		tree.Insert(v)
		fmt.Printf("Добавлен элемент: %d\n", v)
	}

	// Печатаем элементы дерева разными способами
	tree.PrintInOrder()
	tree.PrintPreOrder()
	tree.PrintPostOrder()

	// Поиск элементов
	fmt.Println("\nПоиск элементов:")
	searchValues := []int{40, 25, 70, 90}
	for _, v := range searchValues {
		found := tree.Search(v)
		if found {
			fmt.Printf("Элемент %d найден в дереве\n", v)
		} else {
			fmt.Printf("Элемент %d не найден в дереве\n", v)
		}
	}

	// Удаление элементов
	fmt.Println("\nУдаляем элементы из дерева:")
	deleteValues := []int{20, 30, 50}
	for _, v := range deleteValues {
		fmt.Printf("Удаляем элемент: %d\n", v)
		tree.Delete(v)
		tree.PrintInOrder()
	}
}