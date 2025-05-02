package public

import (
	"domainFilter/utils"
	"fmt"
	"sync"
)

var Version = "1.0"
var Logo = `
  o__ __o                                               o             
 <|     v\                                            _<|>_           
 / \     <\                                                           
 \o/       \o   o__ __o   \o__ __o__ __o     o__ __o/   o   \o__ __o  
  |         |> /v     v\   |     |     |>   /v     |   <|>   |     |> 
 / \       // />       <\ / \   / \   / \  />     / \  / \  / \   / \ 
 \o/      /   \         / \o/   \o/   \o/  \      \o/  \o/  \o/   \o/ 
  |      o     o       o   |     |     |    o      |    |    |     |  
 / \  __/>     <\__ __/>  / \   / \   / \   <\__  / \  / \  / \   / \ 
                                                                      
                                                                      
                                                                      
  o__ __o__/_   o    o   o                                            
 <|    v      _<|>_ <|> <|>                                           
 < >                / \ < >                                           
  |             o   \o/  |       o__  __o  \o__ __o                   
  o__/_        <|>   |   o__/_  /v      |>  |     |>                  
  |            / \  / \  |     />      //  / \   < >                  
 <o>           \o/  \o/  |     \o    o/    \o/                        
  |             |    |   o      v\  /v __o  |        v%s               
 / \           / \  / \  <\__    <\/> __/> / \           by ifacker
`

var (
	Options     *FlagOptions
	Targets     []string
	IpDomainMap map[string][]string = make(map[string][]string) // IP为key，domain为value
	IpMapMutex  sync.Mutex
	FileName    string = fmt.Sprintf("domainFilter_output_%s", utils.GetTimePath())
)
