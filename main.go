package main

import "fmt"

func main() {
	// --- 配列 (Arrays) --- 
	// サイズが3のint型配列を宣言。要素はゼロ値（0）で初期化されます。
	var a1 [3]int
	// サイズが3のint型配列を宣言し、初期値を指定して初期化します。
	var a2 = [3]int{10, 20, 30}
	// 配列のサイズをコンパイラに推論させ、初期値を指定して初期化します。
	a3 := [...]int{10, 20}
	// 各配列の内容を表示します。
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	// a3の長さ（要素数）と容量（確保されているメモリの最大要素数）を表示します。
	fmt.Printf("%v %v\n", len(a3), cap(a3))
	// a2とa3の型を表示します。配列の型は要素の型とサイズによって決まります。
	fmt.Printf("%T %T\n", a2, a3)

	// --- スライス (Slices) --- 
	// nilスライスを宣言。スライスはnil値を持つことができます。
	var s1 []int
	// 空のスライスを宣言。nilではありませんが、長さも容量も0です。
	s2 := []int{}
	// s1の型、値、長さ、容量を表示します。
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// s2の型、値、長さ、容量を表示します。
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	// s1がnilかどうかをチェックします。
	fmt.Println(s1 == nil)
	// s2がnilかどうかをチェックします。
	fmt.Println(s2 == nil)
	// s1に要素を追加します。appendは新しいスライスを返すことがあります。
	s1 = append(s1, 1, 2, 3)
	// 要素追加後のs1の型、値、長さ、容量を表示します。
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// 別のスライスs3を宣言し初期化します。
	s3 := []int{4, 5, 6}
	// s1にs3の要素をすべて追加します（...はスライスを展開します）。
	s1 = append(s1, s3...)
	// s3追加後のs1の型、値、長さ、容量を表示します。
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	// make関数を使って、長さ0、容量2のスライスs4を作成します。
	s4 := make([]int, 0, 2)
	// s4の型、値、長さ、容量を表示します。
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// s4に要素を追加します。容量を超えると新しい基底配列が割り当てられます。
	s4 = append(s4, 1, 2, 3, 4)
	// 要素追加後のs4の型、値、長さ、容量を表示します。
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	// 長さ4、容量6のスライスs5を作成します。
	s5 := make([]int, 4, 6)
	// s5の値、長さ、容量を表示します。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// s5のインデックス1から3（排他的）までのサブスライスs6を作成します。
	s6 := s5[1:3]
	// s6のインデックス1の要素を変更します。これはs5の対応する要素も変更します。
	s6[1] = 10
	// s5の値、長さ、容量を表示します。s6の変更がs5に反映されていることを確認します。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// s6の値、長さ、容量を表示します。
	fmt.Printf("s6: %v %v %v\n", s6, len(s6), cap(s6))
	// s6に要素を追加します。s6の容量がs5の基底配列の残りの容量に依存することに注意してください。
	s6 = append(s6, 2)
	// s5の値、長さ、容量を表示します。s6へのappendがs5に影響を与えないことを確認した場合（容量を超えた場合）。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// append後のs6の値、長さ、容量を表示します。
	fmt.Printf("s6 appended: %v %v %v\n", s6, len(s6), cap(s6))

	// s5のサブスライスと同じ長さの新しいスライスsc6を作成します。
	sc6 := make([]int, len(s5[1:3]))
	// コピー元のs5の値、長さ、容量を表示します。
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	// コピー先のsc6のコピー前の値、長さ、容量を表示します。
	fmt.Printf("sc6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// s5のサブスライスからsc6に要素をコピーします。copyは新しい基底配列を作成します。
	copy(sc6, s5[1:3])
	// コピー後のsc6の値、長さ、容量を表示します。
	fmt.Printf("sc6 dst of copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// sc6の要素を変更します。これはs5には影響しません。
	sc6[1] = 12
	// sc6の変更がs5に影響しないことを確認します。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// sc6の値、長さ、容量を表示します。
	fmt.Printf("sc6: %v %v %v\n", sc6, len(sc6), cap(sc6))

	// 長さ4、容量6のスライスs5を再作成します。
	s5 = make([]int, 4, 6)
	// s5のインデックス1から3（排他的）まで、容量3のサブスライスfs6を作成します。
	// 3番目のインデックスは容量の制限を設定します。
	fs6 := s5[1:3:3]
	// s5の値、長さ、容量を表示します。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// fs6の値、長さ、容量を表示します。
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	// fs6の要素を変更します。
	fs6[0] = 6
	fs6[1] = 7
	// fs6に要素を追加します。容量が3に制限されているため、新しい基底配列が作成されます。
	fs6 = append(fs6, 8)
	// s5の値、長さ、容量を表示します。fs6へのappendがs5に影響を与えないことを確認します。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// append後のfs6の値、長さ、容量を表示します。
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	// s5の要素を変更します。
	s5[3] = 9
	// s5の変更がfs6に影響しないことを確認します。
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// fs6の値、長さ、容量を表示します。
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))

	// --- マップ (Maps) ---
	// nilマップを宣言。マップはnil値を持つことができます。
	var m1 map[string]int
	// 空のマップを宣言。nilではありませんが、要素は0です。
	m2 := map[string]int{}
	// m1の値とnilかどうかを表示します。
	fmt.Printf("%v %v \n", m1, m1 == nil)
	// m2の値とnilかどうかを表示します。
	fmt.Printf("%v %v \n", m2, m2 == nil)
	// マップにキーと値を追加します。
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	// m2の内容、長さ、キー"A"の値を取得して表示します。
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])
	// キー"A"の要素をマップから削除します。
	delete(m2, "A")
	// 削除後のm2の内容、長さ、キー"A"の値（存在しないためゼロ値）を表示します。
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])
	// キー"A"の値と存在チェックの結果を取得して表示します。
	v, ok := m2["A"]
	fmt.Printf("%v %v\n", v, ok)
	// キー"C"の値と存在チェックの結果を取得して表示します。
	v, ok = m2["C"]
	fmt.Printf("%v %v\n", v, ok)

	// m2のすべてのキーと値をループで処理し表示します。
	for k, v := range m2 {
		fmt.Printf("%v %v\n", k, v)
	}
}
