package inflection

import "testing"

func TestCamelize(t *testing.T) {
	type args struct {
		s                    string
		uppercaseFirstLetter bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//testCases["a_a"] = "aA"
		//testCases["aa_a"] = "aaA"
		//testCases["aaa"] = "aaa"
		//testCases["i_love_golang_and_json_so_much"] = "iLoveGolangAndJsonSoMuch"
		{
			name: "camelize without first letter",
			args: args{
				s:                    "",
				uppercaseFirstLetter: false,
			},
			want: "",
		},
		{
			name: "camelize without first letter",
			args: args{
				s:                    "a_a",
				uppercaseFirstLetter: false,
			},
			want: "aA",
		},
		{
			name: "camelize without first letter",
			args: args{
				s:                    "aa_a",
				uppercaseFirstLetter: false,
			},
			want: "aaA",
		},
		{
			name: "camelize without first letter",
			args: args{
				s:                    "aaa",
				uppercaseFirstLetter: false,
			},
			want: "aaa",
		},
		{
			name: "camelize without first letter",
			args: args{
				s:                    "i_love_golang_and_json_so_much",
				uppercaseFirstLetter: false,
			},
			want: "iLoveGolangAndJsonSoMuch",
		},

		//testCases["a_a"] = "AA"
		//testCases["aa_a"] = "AaA"
		//testCases["aa__a"] = "Aa__a"
		//testCases["aaa"] = "Aaa"
		//testCases["i_love_golang_and_json_so_much"] = "ILoveGolangAndJsonSoMuch"
		{
			name: "camelize with first letter",
			args: args{
				s:                    "a_a",
				uppercaseFirstLetter: true,
			},
			want: "AA",
		},
		{
			name: "camelize with first letter",
			args: args{
				s:                    "aa_a",
				uppercaseFirstLetter: true,
			},
			want: "AaA",
		},
		{
			name: "camelize with first letter",
			args: args{
				s:                    "aa__a",
				uppercaseFirstLetter: true,
			},
			want: "Aa__a",
		},
		{
			name: "camelize with first letter",
			args: args{
				s:                    "aaa",
				uppercaseFirstLetter: true,
			},
			want: "Aaa",
		},
		{
			name: "camelize with first letter",
			args: args{
				s:                    "i_love_golang_and_json_so_much",
				uppercaseFirstLetter: true,
			},
			want: "ILoveGolangAndJsonSoMuch",
		},
		{
			name: "camelize number",
			args: args{
				s:                    "1",
				uppercaseFirstLetter: false,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camelize(tt.args.s, tt.args.uppercaseFirstLetter); got != tt.want {
				t.Errorf("Camelize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnderscore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//testCases["aA"] = "a_a"
		//testCases["aaA"] = "aa_a"
		//testCases["aAA"] = "a_aa"
		//testCases["my_case1"] = "my_case1"
		//testCases["ILoveGolangAndJSONSoMuch"] = "i_love_golang_and_json_so_much"
		{
			name: "convert to underscore",
			args: args{s: ""},
			want: "",
		},
		{
			name: "convert to underscore",
			args: args{s: "aA"},
			want: "a_a",
		},
		{
			name: "convert to underscore",
			args: args{s: "aaA"},
			want: "aa_a",
		},
		{
			name: "convert to underscore",
			args: args{s: "aAA"},
			want: "a_aa",
		},
		{
			name: "convert to underscore",
			args: args{s: "my_case1"},
			want: "my_case1",
		},
		{
			name: "convert to underscore",
			args: args{s: "ILoveGolangAndJSONSoMuch"},
			want: "i_love_golang_and_json_so_much",
		},
		{
			name: "convert number",
			args: args{s: "1"},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Underscore(tt.args.s); got != tt.want {
				t.Errorf("Underscore() = %v, want %v", got, tt.want)
			}
		})
	}
}
