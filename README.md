# go-forward

forward 是孤 使用 go 開發的 一個 tcp 隧道 它可以爲 服務器 - 客戶端 之間建立起來 安全的 加密傳輸 並且這些操作 對於服務器 客戶端 來說 是 透明的

#crypto.go

crypto.go 檔案中 定義了兩個 Encryption Decryption 加密/解密 兩個函數 你只需要 簡單的修改 這兩個函數 就可以 替換爲你自己的加密實現 孤使用 king-go 庫的一個加密算法

king-go 是孤另外一個 開源的 go 代碼庫 https://github.com/zuiwuchang/king-go


#forward-s

forward-s 爲隧道服務器 forward-s.json 檔案 定義了 服務器配置

```
{
	"_comment":"服務器工作地址",
	"LAddr":":1102",
	
	"_comment":"客戶端未活動 超時時間 單位 0，如果爲 0 永不超時",
	"Timeout":600,

	"_comment":"數據加密key",
	"Key":"Cerber is an idea",

	"_comment":"驗證密碼",
	"Pwd":"i'm king",

	"_comment":"是否使用 大端序",
	"BigEndian":false,

	"_comment":"要轉接的服務",
	"Service":[
		{
			"_comment":"服務 id 號",
			"Id":1,

			"_comment":"服務 地址",
			"Addr":"127.0.0.1:3389"
		},
		{
			"Id":2,
			"Addr":"127.0.0.1:9666"
		}
	],
	"_comment":"要打印的日誌 信息 TRACE DEBUG INFO ERROR FAULT",
	"Logs":[
		"TRACE",
		"DEBUG",
		"INFO",
		"ERROR",
		"FAULT"
	],
	"_comment":"日誌是否要打印 代碼行",
	"LogLine":true
}
```

上面 的配置 在服務器中 創建了 兩個 隧道 id 分別爲 1 2 隧道的地址 分別 爲 127.0.0.1:3389 127.0.0.1:9666

此時 如果 客戶端 連接 forward-s 並且 請求 服務 1 則會 建立一個 到服務器 127.0.0.1:3389 的 安全 隧道


#forward-c

forward-c 會爲 隧道 客戶端 forward-c 會打開指定 的端口 並且 將此端口 和 服務器 之間 建立一個 安全的 隧道 forward-c.json 定義了 forward-c 的行爲

```
{
	"_comment":"本地監聽地址",
	"LAddr":":6666",
	
	"_comment":"服務器工作地址",
	"Addr1":"127.0.0.1:1102",
	"Addr":"192.168.1.204:3000",

	"_comment":"數據加密key",
	"Key":"Cerber is an idea",

	"_comment":"驗證密碼",
	"Pwd":"i'm king",

	"_comment":"是否使用 大端序",
	"BigEndian":false,

	"_comment":"要轉接的服務id",
	"Service":1,

	"_comment":"要打印的日誌 信息 TRACE DEBUG INFO ERROR FAULT",
	"Logs":[
		"TRACE",
		"DEBUG",
		"INFO",
		"ERROR",
		"FAULT"
	],
	"_comment":"日誌是否要打印 代碼行",
	"LogLine":true
}
```
上面的 配置 將開啓 :6666 端口 並且 將其 和 服務器 的 服務 1 建立 安全 連接

配合 forward-s 的配置 看 此時 只需要 使用 3389 客戶端 連接 本機的 6666 端口 就可以安全的連接到服務器遠程桌面
