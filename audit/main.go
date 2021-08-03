package main

import (
	"fmt"
	"student"
)

func main() {

	// // ====> Atoi <====
	// fmt.Println("====> Atoi <====")
	// fmt.Println(student.Atoi(""))
	// fmt.Println(student.Atoi("-"))
	// fmt.Println(student.Atoi("--123"))
	// fmt.Println(student.Atoi("1"))
	// fmt.Println(student.Atoi("-3"))
	// fmt.Println(student.Atoi("8292"))
	// fmt.Println(student.Atoi("9223372036854775807"))
	// fmt.Println(student.Atoi("-9223372036854775808"))

	// // ====> RecursivePower <====
	// fmt.Println("\n====> RecursivePower <====")
	// fmt.Println(student.RecursivePower(-7, -2))
	// fmt.Println(student.RecursivePower(-8, -7))
	// fmt.Println(student.RecursivePower(4, 8))
	// fmt.Println(student.RecursivePower(1, 3))
	// fmt.Println(student.RecursivePower(-1, 1))
	// fmt.Println(student.RecursivePower(-6, 5))

	// // ====> PrintCombN <====
	// fmt.Println("\n====> PrintCombN <====")
	// fmt.Println("1----------------------------")
	// student.PrintCombN(1)
	// fmt.Println("2----------------------------")
	// student.PrintCombN(2)
	// fmt.Println("3----------------------------")
	// student.PrintCombN(3)
	// fmt.Println("4----------------------------")
	// student.PrintCombN(4)
	// fmt.Println("5----------------------------")
	// student.PrintCombN(5)
	// fmt.Println("6----------------------------")
	// student.PrintCombN(6)
	// fmt.Println("7----------------------------")
	// student.PrintCombN(7)
	// fmt.Println("8----------------------------")
	// student.PrintCombN(8)
	// fmt.Println("9----------------------------")
	// student.PrintCombN(9)

	// // ====> PrintNbrBase <====
	// fmt.Println("\n====> PrintNbrBase <====")
	// student.PrintNbrBase(919617, "01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-661165, "1")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-861737, "Zone01Zone01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "0123456789ABCDEF")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-125, "choumi")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "-ab")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-9223372036854775808, "0123456789")
	// z01.PrintRune('\n')

	// // ====> AtoiBase <====
	// fmt.Println("\n====> AtoiBase <====")
	// fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	// fmt.Println(student.AtoiBase("0001", "01"))
	// fmt.Println(student.AtoiBase("00", "01"))
	// fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	// fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	// fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
	// fmt.Println(student.AtoiBase("125", "0123456789"))
	// fmt.Println(student.AtoiBase("uoi", "choumi"))
	// fmt.Println(student.AtoiBase("bbbbbab", "-ab"))

	// ====> SplitWhiteSpaces <====
	// fmt.Println("\n====> SplitWhiteSpaces <====")
	// fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	// fmt.Println(student.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"))
	// fmt.Println(student.SplitWhiteSpaces("aiding in computations such as multiplication and division ."))
	// fmt.Println(student.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
	// fmt.Println(student.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
	// fmt.Println(student.SplitWhiteSpaces("In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))
	fmt.Println(student.SplitWhiteSpaces("        asad\n     "))

	// ====> Split <====
	// fmt.Println("\n====> Split <====")
	// fmt.Println(student.Split("|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,", "|=choumi=|"))
	// fmt.Println(student.Split("!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,", "!==!"))
	// fmt.Println(student.Split("AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,", "AFJ"))

	fmt.Println(student.Split("", "aaa"))

	// 	// ====> ConvertBase <====
	// 	fmt.Println("\n====> ConvertBase <====")
	// 	fmt.Println(student.ConvertBase("4506C", "0123456789ABCDEF", "choumi"))
	// 	fmt.Println(student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF"))
	// 	fmt.Println(student.ConvertBase("256850", "0123456789", "01"))
	// 	fmt.Println(student.ConvertBase("uuhoumo", "choumi", "Zone01"))
	// 	fmt.Println(student.ConvertBase("683241", "0123456789", "0123456789"))

	// 	// ====> AdvancedSortWordArr <====
	// 	fmt.Println("\n====> AdvancedSortWordArr <====")
	// 	slice := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	// 	student.AdvancedSortWordArr(slice, strings.Compare)
	// 	fmt.Println(slice)
	// 	slice = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	// 	student.AdvancedSortWordArr(slice, strings.Compare)
	// 	fmt.Println(slice)
	// 	slice = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// 	student.AdvancedSortWordArr(slice, strings.Compare)
	// 	fmt.Println(slice)
	// 	slice = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// 	student.AdvancedSortWordArr(slice, student.Compare)
	// 	fmt.Println(slice)
	// 	slice = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	// 	student.AdvancedSortWordArr(slice, student.Compare)
	// 	fmt.Println(slice)
	// 	slice = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	// 	student.AdvancedSortWordArr(slice, student.Compare)
	// 	fmt.Println(slice)

	// 	// ====> ActiveBits <====
	// 	fmt.Println("\n====> ActiveBits <====")
	// 	fmt.Println(student.ActiveBits(15))
	// 	fmt.Println(student.ActiveBits(17))
	// 	fmt.Println(student.ActiveBits(4))
	// 	fmt.Println(student.ActiveBits(11))
	// 	fmt.Println(student.ActiveBits(9))
	// 	fmt.Println(student.ActiveBits(12))
	// 	fmt.Println(student.ActiveBits(2))

	// 	// ====> SortListInsert <====
	// 	fmt.Println("\n====> SortListInsert <====")
	// 	var link1 *student.NodeI
	// 	link1 = listPushBack(link1, 0)
	// 	link1 = student.SortListInsert(link1, 39)
	// 	PrintList(link1)
	// 	var link2 *student.NodeI
	// 	link2 = listPushBack(link2, 0)
	// 	link2 = listPushBack(link2, 1)
	// 	link2 = listPushBack(link2, 2)
	// 	link2 = listPushBack(link2, 3)
	// 	link2 = listPushBack(link2, 4)
	// 	link2 = listPushBack(link2, 5)
	// 	link2 = listPushBack(link2, 24)
	// 	link2 = listPushBack(link2, 25)
	// 	link2 = listPushBack(link2, 54)
	// 	link2 = student.SortListInsert(link2, 33)
	// 	PrintList(link2)
	// 	var link3 *student.NodeI
	// 	link3 = listPushBack(link3, 0)
	// 	link3 = listPushBack(link3, 2)
	// 	link3 = listPushBack(link3, 18)
	// 	link3 = listPushBack(link3, 33)
	// 	link3 = listPushBack(link3, 37)
	// 	link3 = listPushBack(link3, 37)
	// 	link3 = listPushBack(link3, 39)
	// 	link3 = listPushBack(link3, 52)
	// 	link3 = listPushBack(link3, 53)
	// 	link3 = listPushBack(link3, 57)
	// 	link3 = student.SortListInsert(link3, 53)
	// 	PrintList(link3)
	// 	var link4 *student.NodeI
	// 	link4 = listPushBack(link4, 0)
	// 	link4 = listPushBack(link4, 5)
	// 	link4 = listPushBack(link4, 18)
	// 	link4 = listPushBack(link4, 24)
	// 	link4 = listPushBack(link4, 28)
	// 	link4 = listPushBack(link4, 35)
	// 	link4 = listPushBack(link4, 42)
	// 	link4 = listPushBack(link4, 45)
	// 	link4 = listPushBack(link4, 52)
	// 	link4 = student.SortListInsert(link4, 52)
	// 	PrintList(link4)
	// 	var link5 *student.NodeI
	// 	link5 = listPushBack(link5, 0)
	// 	link5 = listPushBack(link5, 12)
	// 	link5 = listPushBack(link5, 20)
	// 	link5 = listPushBack(link5, 23)
	// 	link5 = listPushBack(link5, 23)
	// 	link5 = listPushBack(link5, 24)
	// 	link5 = listPushBack(link5, 30)
	// 	link5 = listPushBack(link5, 41)
	// 	link5 = listPushBack(link5, 53)
	// 	link5 = listPushBack(link5, 57)
	// 	link5 = listPushBack(link5, 59)
	// 	link5 = student.SortListInsert(link5, 38)
	// 	PrintList(link5)

	// 	// ====> SortedListMerge <====
	// 	fmt.Println("\n====> SortedListMerge <====")
	// 	srtListMrg()

	// 	// ====> ListRemoveIf <====
	// 	fmt.Println("\n====> ListRemoveIf <====")
	// 	lremoveif()

	// 	// ====> BTreeTransplant <====
	// 	fmt.Println("\n====> BTreeTransplant <====")
	// 	root := &student.TreeNode{Data: "01"}
	// 	student.BTreeInsertData(root, "07")
	// 	student.BTreeInsertData(root, "05")
	// 	student.BTreeInsertData(root, "12")
	// 	student.BTreeInsertData(root, "10")
	// 	student.BTreeInsertData(root, "02")
	// 	student.BTreeInsertData(root, "03")
	// 	node := student.BTreeSearchItem(root, "12")
	// 	replacement := &student.TreeNode{Data: "55"}
	// 	student.BTreeInsertData(replacement, "60")
	// 	student.BTreeInsertData(replacement, "33")
	// 	student.BTreeInsertData(replacement, "12")
	// 	student.BTreeInsertData(replacement, "15")
	// 	root = student.BTreeTransplant(root, node, replacement)
	// 	student.BtreePrint(root)

	// 	root = &student.TreeNode{Data: "33"}
	// 	student.BTreeInsertData(root, "5")
	// 	student.BTreeInsertData(root, "20")
	// 	student.BTreeInsertData(root, "52")
	// 	student.BTreeInsertData(root, "31")
	// 	student.BTreeInsertData(root, "13")
	// 	student.BTreeInsertData(root, "11")
	// 	node = student.BTreeSearchItem(root, "20")
	// 	replacement = &student.TreeNode{Data: "55"}
	// 	student.BTreeInsertData(replacement, "60")
	// 	student.BTreeInsertData(replacement, "33")
	// 	student.BTreeInsertData(replacement, "12")
	// 	student.BTreeInsertData(replacement, "15")
	// 	root = student.BTreeTransplant(root, node, replacement)
	// 	student.BtreePrint(root)

	// 	root = &student.TreeNode{Data: "03"}
	// 	student.BTreeInsertData(root, "39")
	// 	student.BTreeInsertData(root, "99")
	// 	student.BTreeInsertData(root, "44")
	// 	student.BTreeInsertData(root, "11")
	// 	student.BTreeInsertData(root, "14")
	// 	student.BTreeInsertData(root, "11")
	// 	node = student.BTreeSearchItem(root, "11")
	// 	replacement = &student.TreeNode{Data: "55"}
	// 	student.BTreeInsertData(replacement, "60")
	// 	student.BTreeInsertData(replacement, "33")
	// 	student.BTreeInsertData(replacement, "12")
	// 	student.BTreeInsertData(replacement, "15")
	// 	root = student.BTreeTransplant(root, node, replacement)
	// 	student.BtreePrint(root)

	// 	// ====> BTreeApplyByLevel <====
	// 	fmt.Println("\n====> BTreeApplyByLevel <====")
	// 	root = &student.TreeNode{Data: "01"}
	// 	student.BTreeInsertData(root, "07")
	// 	student.BTreeInsertData(root, "12")
	// 	student.BTreeInsertData(root, "10")
	// 	student.BTreeInsertData(root, "05")
	// 	student.BTreeInsertData(root, "02")
	// 	student.BTreeInsertData(root, "03")
	// 	student.BTreeApplyByLevel(root, fmt.Println)
	// 	fmt.Println()
	// 	root = &student.TreeNode{Data: "01"}
	// 	student.BTreeInsertData(root, "07")
	// 	student.BTreeInsertData(root, "12")
	// 	student.BTreeInsertData(root, "10")
	// 	student.BTreeInsertData(root, "05")
	// 	student.BTreeInsertData(root, "02")
	// 	student.BTreeInsertData(root, "03")
	// 	student.BTreeApplyByLevel(root, fmt.Print)
	// 	fmt.Println()

	// 	// ====> BTreeDeleteNode <====
	// 	fmt.Println("\n====> BTreeDeleteNode <====")
	// 	root = &student.TreeNode{Data: "01"}
	// 	student.BTreeInsertData(root, "07")
	// 	student.BTreeInsertData(root, "12")
	// 	student.BTreeInsertData(root, "10")
	// 	student.BTreeInsertData(root, "05")
	// 	student.BTreeInsertData(root, "02")
	// 	student.BTreeInsertData(root, "03")
	// 	nodel := student.BTreeSearchItem(root, "02")
	// 	student.BtreePrint(student.BTreeDeleteNode(root, nodel))

	// 	root = &student.TreeNode{Data: "33"}
	// 	student.BTreeInsertData(root, "5")
	// 	student.BTreeInsertData(root, "20")
	// 	student.BTreeInsertData(root, "52")
	// 	student.BTreeInsertData(root, "31")
	// 	student.BTreeInsertData(root, "13")
	// 	student.BTreeInsertData(root, "11")
	// 	nodel = student.BTreeSearchItem(root, "20")
	// 	student.BtreePrint(student.BTreeDeleteNode(root, nodel))

	// 	root = &student.TreeNode{Data: "03"}
	// 	student.BTreeInsertData(root, "39")
	// 	student.BTreeInsertData(root, "11")
	// 	student.BTreeInsertData(root, "99")
	// 	student.BTreeInsertData(root, "14")
	// 	student.BTreeInsertData(root, "44")
	// 	student.BTreeInsertData(root, "11")
	// 	nodel = student.BTreeSearchItem(root, "03")
	// 	student.BtreePrint(student.BTreeDeleteNode(root, nodel))

	// 	root = &student.TreeNode{Data: "03"}
	// 	student.BTreeInsertData(root, "03")
	// 	student.BTreeInsertData(root, "01")
	// 	student.BTreeInsertData(root, "94")
	// 	student.BTreeInsertData(root, "19")
	// 	student.BTreeInsertData(root, "24")
	// 	student.BTreeInsertData(root, "111")
	// 	nodel = student.BTreeSearchItem(root, "03")
	// 	student.BtreePrint(student.BTreeDeleteNode(root, nodel))
	// }

	// func PrintList(l *student.NodeI) {
	// 	it := l
	// 	for it != nil {
	// 		fmt.Print(it.Data, " -> ")
	// 		it = it.Next
	// 	}
	// 	fmt.Print(nil, "\n")
	// }

	// func listPushBack(l *student.NodeI, data int) *student.NodeI {
	// 	n := &student.NodeI{Data: data}

	// 	if l == nil {
	// 		return n
	// 	}
	// 	iterator := l
	// 	for iterator.Next != nil {
	// 		iterator = iterator.Next
	// 	}
	// 	iterator.Next = n
	// 	return l
	// }

	// func srtListMrg() {
	// 	var link *student.NodeI
	// 	var link2 *student.NodeI
	// 	PrintList(student.SortedListMerge(link, link2))
	// 	link, link2 = nil, nil
	// 	link2 = listPushBack(link2, 2)
	// 	link2 = listPushBack(link2, 2)
	// 	link2 = listPushBack(link2, 4)
	// 	link2 = listPushBack(link2, 9)
	// 	link2 = listPushBack(link2, 12)
	// 	link2 = listPushBack(link2, 12)
	// 	link2 = listPushBack(link2, 19)
	// 	link2 = listPushBack(link2, 20)
	// 	PrintList(student.SortedListMerge(link, link2))
	// 	link, link2 = nil, nil

	// 	link = listPushBack(link, 4)
	// 	link = listPushBack(link, 4)
	// 	link = listPushBack(link, 6)
	// 	link = listPushBack(link, 9)
	// 	link = listPushBack(link, 13)
	// 	link = listPushBack(link, 18)
	// 	link = listPushBack(link, 20)
	// 	link = listPushBack(link, 20)
	// 	PrintList(student.SortedListMerge(link, link2))

	// 	link, link2 = nil, nil

	// 	link = listPushBack(link, 0)
	// 	link = listPushBack(link, 7)
	// 	link = listPushBack(link, 39)
	// 	link = listPushBack(link, 92)
	// 	link = listPushBack(link, 97)
	// 	link = listPushBack(link, 93)
	// 	link = listPushBack(link, 91)
	// 	link = listPushBack(link, 28)
	// 	link = listPushBack(link, 64)

	// 	link2 = listPushBack(link2, 80)
	// 	link2 = listPushBack(link2, 23)
	// 	link2 = listPushBack(link2, 27)
	// 	link2 = listPushBack(link2, 30)
	// 	link2 = listPushBack(link2, 85)
	// 	link2 = listPushBack(link2, 81)
	// 	link2 = listPushBack(link2, 75)
	// 	link2 = listPushBack(link2, 70)
	// 	PrintList(student.SortedListMerge(link, link2))

	// 	link, link2 = nil, nil

	// 	link = listPushBack(link, 0)
	// 	link = listPushBack(link, 2)
	// 	link = listPushBack(link, 11)
	// 	link = listPushBack(link, 30)
	// 	link = listPushBack(link, 54)
	// 	link = listPushBack(link, 56)
	// 	link = listPushBack(link, 70)
	// 	link = listPushBack(link, 79)
	// 	link = listPushBack(link, 99)

	// 	link2 = listPushBack(link2, 23)
	// 	link2 = listPushBack(link2, 28)
	// 	link2 = listPushBack(link2, 38)
	// 	link2 = listPushBack(link2, 67)
	// 	link2 = listPushBack(link2, 67)
	// 	link2 = listPushBack(link2, 79)
	// 	link2 = listPushBack(link2, 95)
	// 	link2 = listPushBack(link2, 97)
	// 	PrintList(student.SortedListMerge(link, link2))

	// 	link, link2 = nil, nil

	// 	link = listPushBack(link, 0)
	// 	link = listPushBack(link, 3)
	// 	link = listPushBack(link, 8)
	// 	link = listPushBack(link, 8)
	// 	link = listPushBack(link, 13)
	// 	link = listPushBack(link, 19)
	// 	link = listPushBack(link, 34)
	// 	link = listPushBack(link, 38)
	// 	link = listPushBack(link, 46)

	// 	link2 = listPushBack(link2, 7)
	// 	link2 = listPushBack(link2, 39)
	// 	link2 = listPushBack(link2, 45)
	// 	link2 = listPushBack(link2, 53)
	// 	link2 = listPushBack(link2, 59)
	// 	link2 = listPushBack(link2, 70)
	// 	link2 = listPushBack(link2, 76)
	// 	link2 = listPushBack(link2, 79)
	// 	PrintList(student.SortedListMerge(link, link2))
	// }

	// func lremoveif() {
	// 	link := &student.List{}
	// 	student.ListRemoveIf(link, 1)
	// 	PrintListL(link)
	// 	student.ListRemoveIf(link, 96)
	// 	PrintListL(link)
	// 	student.ListPushBack(link, 98)
	// 	student.ListPushBack(link, 98)
	// 	student.ListPushBack(link, 33)
	// 	student.ListPushBack(link, 34)
	// 	student.ListPushBack(link, 33)
	// 	student.ListPushBack(link, 34)
	// 	student.ListPushBack(link, 33)
	// 	student.ListPushBack(link, 89)
	// 	student.ListPushBack(link, 33)
	// 	student.ListRemoveIf(link, 34)
	// 	PrintListL(link)

	// 	link = &student.List{}
	// 	student.ListPushBack(link, 79)
	// 	student.ListPushBack(link, 74)
	// 	student.ListPushBack(link, 99)
	// 	student.ListPushBack(link, 79)
	// 	student.ListPushBack(link, 7)
	// 	student.ListRemoveIf(link, 99)
	// 	PrintListL(link)

	// 	link = &student.List{}
	// 	student.ListPushBack(link, 56)
	// 	student.ListPushBack(link, 93)
	// 	student.ListPushBack(link, 68)
	// 	student.ListPushBack(link, 56)
	// 	student.ListPushBack(link, 87)
	// 	student.ListPushBack(link, 68)
	// 	student.ListPushBack(link, 56)
	// 	student.ListPushBack(link, 68)
	// 	student.ListRemoveIf(link, 68)
	// 	PrintListL(link)

	// 	link = &student.List{}
	// 	student.ListPushBack(link, "mvkUxbqhQve4l")
	// 	student.ListPushBack(link, "4Zc4t hnf SQ")
	// 	student.ListPushBack(link, "q2If E8BPuX ")
	// 	student.ListRemoveIf(link, "4Zc4t hnf SQ")
	// 	PrintListL(link)

}

// func PrintListL(l *student.List) {
// 	if l == nil {
// 		fmt.Println(nil)
// 		return
// 	}
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}

// 	fmt.Print(nil, "\n")
// }
