package rx

// func Test_New_NoGo(t *testing.T) {
// 	r := require.New(t)
//
// 	bb := &bytes.Buffer{}
//
// 	g, err := New(&Options{
// 		Out: NewWriter(bb),
// 	})
// 	r.NoError(err)
//
// 	run := gentest.NewRunner()
// 	run.With(g)
//
// 	run.LookPathFn = func(s string) (string, error) {
// 		return "", exec.ErrNotFound
// 	}
//
// 	r.NoError(run.Run())
//
// 	r.Contains(bb.String(), "executable could not")
// }
