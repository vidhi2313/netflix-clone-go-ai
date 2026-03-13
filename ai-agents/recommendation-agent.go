
package agents

func RecommendMovies(history []string) []string {

    rec := []string{}

    for _, h := range history {

        if h == "Action" {
            rec = append(rec, "John Wick", "Extraction")
        }

        if h == "SciFi" {
            rec = append(rec, "Interstellar", "Dune")
        }
    }

    return rec
}// code ends here 
