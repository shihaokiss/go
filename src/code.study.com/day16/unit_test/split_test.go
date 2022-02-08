package strings_split

import (
	// "reflect"
	"testing")


// func Test_Split(t *testing.T) {
// 	get := Split("ab:cd:ef", ":")
// 	want := []string{"ab", "cd", "ef"}
//     if !reflect.DeepEqual(get, want) {
//         t.Fatalf("get:%v, want:%v", get, want)
// 	}
// }

// func Test2Split(t *testing.T) {
// 	get := Split(":::", ":")
// 	want := []string{"", "", "", ""}
//     if !reflect.DeepEqual(get, want) {
//         t.Fatalf("get:%v, want:%v", get, want)
// 	}
// }
// func Test3Split(t *testing.T) {
// 	get := Split("ababccabab", "cc")
// 	want := []string{"abab", "abab"}
//     if !reflect.DeepEqual(get, want) {
//         t.Fatalf("get:%v, want:%v", get, want)
// 	}
// }

// func TestGroup(t *testing.T) {
// 	type test struct {
// 		input string
// 		seq string
// 		want []string
// 	}
//     tests := map[string]test {
// 		"case1": {input: "aaabccc", seq: "b", want: []string{"aaa", "ccc"}},
// 		"case2": {input: "aaabccc", seq: "bc", want: []string{"aaa", "cc"}},
// 		"case3": {input: "aaabccc", seq: "bcc", want: []string{"aaa", "c"}},
// 	}

// 	for name, tc := range tests { 
// 		t.Run(name, func(t *testing.T){
// 			get := Split(tc.input, tc.seq)
// 			want := tc.want
// 			if !reflect.DeepEqual(get, want) {
// 				t.Fatalf("name: %s, get: %#v, want: %#v", name, get, want)
// 			}
// 		})
// 	}
// }




// 基准测试
func BenchmarkSplit1(b *testing.B) {
    for i:=0; i<b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
