     1	     1	package main
     2	     2	
     3	     3	import (
     4	     4		"encoding/json"
     5	     5		"fmt"
     6	     6		"net/http"
     7	     7		"strconv"
     8	     8		"strings"
     9	     9		"sync"
    10	    10	)
    11	    11	
    12	    12	type Task struct {
    13	    13		ID        int    `json:"id"`
    14	    14		Title     string `json:"title"`
    15	    15		Completed bool   `json:"completed"`
    16	    16	}
    17	    17	
    18	    18	var (
    19	    19		mu      sync.Mutex
    20	    20		tasks   = []Task{
    21			{ID: 1, Title: "Learn Go", Completed: false},
    22			{ID: 2, Title: "Build a REST API", Completed: false},
    23		}
    24	    21		nextID  = 3
    25	    22	)
    26	    23	
    27	    24	func main() {
    28	    25		http.HandleFunc("/tasks", handleTasks)
    29	    26		http.HandleFunc("/tasks/", handleTaskByID)
    30	    27		fmt.Println("taskify-api listening on :8080")
    31	    28		http.ListenAndServe(":8080", nil)
    32	    29	}
    33	    30	
    34	    31	func handleTasks(w http.ResponseWriter, r *http.Request) {
    35	    32		switch r.Method {
    36	    33		case http.MethodGet:
    37	    34			mu.Lock()
    38	    35			defer mu.Unlock()
    39	    36			w.Header().Set("Content-Type", "application/json")
    40	    37			json.NewEncoder(w).Encode(tasks)
    41	    38		case http.MethodPost:
    42	    39			var t Task
    43	    40			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
    44	    41				http.Error(w, "bad request", http.StatusBadRequest)
    45	    42				return
    46	    43			}
    47	    44			mu.Lock()
    48	    45			t.ID = nextID
    49	    46			nextID++
    50	    47			tasks = append(tasks, t)
    51	    48			mu.Unlock()
    52	    49			w.Header().Set("Content-Type", "application/json")
    53	    50			w.WriteHeader(http.StatusCreated)
    54	    51			json.NewEncoder(w).Encode(t)
    55	    52		default:
    56	    53			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    57	    54		}
    58	    55	}
    59	    56	
    60	    57	func handleTaskByID(w http.ResponseWriter, r *http.Request) {
    61	    58		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
    62	    59		if len(parts) < 1 {
    63	    60			http.Error(w, "not found", http.StatusNotFound)
    64	    61			return
    65	    62		}
    66	    63		id, err := strconv.Atoi(parts[0])
    67	    64		if err != nil {
    68	    65			http.Error(w, "invalid id", http.StatusBadRequest)
    69	    66			return
    70	    67		}
    71	    68		if len(parts) == 2 && parts[1] == "complete" && r.Method == http.MethodPatch {
    72	    69			mu.Lock()
    73	    70			defer mu.Unlock()
    74	    71			for i, t := range tasks {
    75	    72				if t.ID == id {
    76	    73					tasks[i].Completed = true
    77	    74					w.Header().Set("Content-Type", "application/json")
    78	    75					json.NewEncoder(w).Encode(tasks[i])
    79	    76					return
    80	    77				}
    81	    78			}
    82	    79			http.Error(w, "not found", http.StatusNotFound)
    83	    80			return
    84	    81		}
    85	    82		http.Error(w, "not found", http.StatusNotFound)
    86	    83	}
