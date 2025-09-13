package main

import (
    "fmt"
	"sync"
	"net/http"
)

func CheckURL(wg *sync.WaitGroup, jobs <-chan string, ){
	defer wg.Done()
	for URL := range jobs{
		resp, err := http.Get(URL)
		if err != nil {
			fmt.Printf("Ошибка запроса к странице %s: %v\n", URL, err)
			return
		} 
		fmt.Printf("URL: %s, Статус: %d\n", URL, resp.StatusCode)
		resp.Body.Close()
	}
}


func main(){
	var wg sync.WaitGroup
	jobs := make(chan string, 9)
    URLs := []string{
        "https://ru.wikipedia.org/wiki/Вики",
        "https://yandex.ru/pogoda/ru/kurgan",
        "https://ru.pinterest.com/ideas",
        "http://google.com", 
        "https://www.timeanddate.com/timer",
    }
	for w := 1; w <= 3; w++ {
		wg.Add(1)
    	go CheckURL(&wg, jobs)
    }
	for _, URL := range URLs{
		jobs <- URL
	}
	close(jobs) 
	wg.Wait() 
}
