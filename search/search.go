package search
import(
	"log"
	"sync"
)
var matchers=make(map[string]Matcher)
func Run(searchTerm string){
	//获取需要搜索的数据源列表
	feeds,err:=RetrieveFeeds()
	if err!=nil{
		log.Fatal(err)
	}
	//无缓冲通道
	results:=make(chan *Result)
	var wg sync.WaitGroup
	wg.Add(len(feed))
	//为每个数据源启动一个goroutine来查找结果
	for _,feed:=range feeds{

		matcher,exists:=matchers[feed.type]
		if !exists{
			matcher=matchers["default"]
		}
		go func(matcher Matchere,feed *Feed) {
			Match(matcher,feed,searchTerm,results)
			wg.Done()
		}(matcher,feed)

	}
	go func(){
		wg.Wait()
		close(results)
	}()
	Display(results)

}