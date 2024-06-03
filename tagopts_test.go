package tagopts

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestPartseTag(t *testing.T) {
	t.Log("Testing ParseTag.")
	{
		testID := 0
		t.Logf("\tTest %2d: ParseTag(\"\", \"\").", testID)
		{
			pv, po := ParseTag("", "")

			if pv != "" {
				t.Fatalf("\t\t%s Invalid result value %q, expected \"\".", failed, pv)
				return
			}
			if len(po) != 0 {
				t.Fatalf("\t\t%s Invalid result options \"%v\", expected [].", failed, po)
				return
			}
			if po.Count() != 0 {
				t.Fatalf("\t\t%s Invalid options count %d, expected 0.", failed, po.Count())
				return
			}

			t.Logf("\t\t%s Parsed successfully.", success)
		}

		testID++
		t.Logf("\tTest %2d: ParseTag(\"-\",\"\").", testID)
		{
			pv, po := ParseTag("-", "")

			if pv != "-" {
				t.Fatalf("\t\t%s Invalid result value %q, expected \"-\".", failed, pv)
				return
			}
			if len(po) != 0 {
				t.Fatalf("\t\t%s Invalid result options %v, expected [].", failed, po)
				return
			}
			if po.Count() != 0 {
				t.Fatalf("\t\t%s Invalid options count %d, expected 0.", failed, po.Count())
				return
			}

			t.Logf("\t\t%s Parsed successfully.", success)
		}

		testID++
		t.Logf("\tTest %2d: ParseTag(\"field,omitifempty\",\"\").", testID)
		{
			pv, po := ParseTag("field,omitifempty", "")

			if pv != "field" {
				t.Fatalf("\t\t%s Invalid result value %q, expected \"field\".", failed, pv)
				return
			}
			if len(po) != 1 {
				t.Fatalf("\t\t%s Invalid result options %v, expected [omitifempty].", failed, po)
				return
			}
			if po.Count() != 1 {
				t.Fatalf("\t\t%s Invalid options count %d, expected 1.", failed, po.Count())
				return
			}
			if !po.Contains("omitifempty") {
				t.Fatalf("\t\t%s Options is not contained %q.", failed, "omitifempty")
				return
			}
			if po.Contains("nonexistant") {
				t.Fatalf("\t\t%s Options is contained %q.", failed, "nonexistant")
				return
			}
			if po.Option(0) != "omitifempty" {
				t.Fatalf("\t\t%s Invalid options.Option(0)  returns %q, expected \"omitifempty\".", failed, po.Option(0))
				return
			}
			if po.Option(1) != "" {
				t.Fatalf("\t\t%s Invalid options.Option(1)  returns %q, expected \"\".", failed, po.Option(1))
				return
			}

			t.Logf("\t\t%s Parsed successfully.", success)
		}

		testID++
		t.Logf("\t\tTest %2d: ParseTag(\"f,opt1,opt2\",\"\").", testID)
		{
			pv, po := ParseTag("f,opt1,opt2", "")

			if pv != "f" {
				t.Fatalf("\t\t%s Invalid result value %q, expected \"f\".", failed, pv)
				return
			}
			if len(po) != 2 {
				t.Fatalf("\t\t%s Invalid result options %v, expected [opt1,opt2].", failed, po)
				return
			}
			if po.Count() != 2 {
				t.Fatalf("\t\t%s Invalid options count %d, expected 2.", failed, po.Count())
				return
			}
			if !po.Contains("opt1") {
				t.Fatalf("\t\t%s Options is not contained %q.", failed, "opt1")
				return
			}
			if !po.Contains("opt2") {
				t.Fatalf("\t\t%s Options is not contained %q.", failed, "opt2")
				return
			}
			if po.Contains("nonexistant") {
				t.Fatalf("\t\t%s Options is contained %s.", failed, "nonexistant")
				return
			}
			if po.Option(0) != "opt1" {
				t.Fatalf("\t\t%s Invalid options.Option(0) returns %q, expected \"opt1\".", failed, po.Option(0))
				return
			}
			if po.Option(1) != "opt2" {
				t.Fatalf("\t\t%s Invalid options.Option(1) returns %q, expected \"opt2\".", failed, po.Option(1))
				return
			}
			if po.Option(2) != "" {
				t.Fatalf("\t\t%s Invalid options.Option(2) returns %q, expected \"\".", failed, po.Option(2))
				return
			}

			t.Logf("\t\t%s Parsed successfully.", success)
		}

		testID++
		t.Logf("\tTest %2d: ParseTag(\",opt\",\"default\").", testID)
		{
			pv, po := ParseTag(",opt", "default")

			if pv != "default" {
				t.Fatalf("\t\t%s Invalid result value %q, expected \"default\".", failed, pv)
				return
			}
			if len(po) != 1 {
				t.Fatalf("\t\t%s Invalid result options %v, expected [opt].", failed, po)
				return
			}
			if po.Count() != 1 {
				t.Fatalf("\t%s\tInvalid options count \"%d\", expected 1.", failed, po.Count())
				return
			}
			if !po.Contains("opt") {
				t.Fatalf("\t\t%s Options is not contained %q.", failed, "opt")
				return
			}
			if po.Contains("nonexistant") {
				t.Fatalf("\t\t%s Options is contained %q.", failed, "nonexistant")
				return
			}
			if po.Option(0) != "opt" {
				t.Fatalf("\t\t%s Invalid options.Option(0) returns %q, expected \"opt\".", failed, po.Option(0))
				return
			}
			if po.Option(1) != "" {
				t.Fatalf("\t\t%s Invalid options.Option(1) returns %q, expected \"\".", failed, po.Option(1))
				return
			}

			t.Logf("\t\t%s Parsed successfully.", success)
		}
	}
}
