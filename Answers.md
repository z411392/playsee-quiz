### infraDesign

> 假設公司有個訂單系統, 所有使用者的資料都存在MySQL資料庫底下, 請設計一個報表系統提供分析數據, 
> 並滿足「在資料量大的情況下系統仍可以正常運作」
> 試著描述出此系統的大致架構與思維

1. 一定要做 partition，partition 可能是用訂單建立時間的「年+月」（逐月）。
2. 因為 limit + offset 還是要把大量的資料都跑過，分頁時應該要改用 `where orders.created_at >= 開始時間 and orders.created_at < 結束時間`。
3. MySQL 的設定方面要提升允許的最大 Connections 數，縮短關閉、回收 Connection 的時間。 -> 但我沒設定過。
4. 資料庫的主機考慮 Single Master + Multi Slaves 的架構。 -> 但我沒實作過。
5. 一定要能夠備份、救援⋯⋯。-> 但我沒實作過。

### NoSQL
#### 問題A

> 設計一個NoSQL DB的rowkey，並說明設計原因，滿足
>   - 找出某個user的post
>   - 可由新到舊且由舊到新查找
>   - 依照NoSQL DB特性，避免hotspot產生

猜測這裡用的 NoSQL 是 HBase，我之前沒用過。

按照網路上說法，RowKey 跟 MySQL 吃索引的原則一樣，都是從前面比對至後面。

如果要找出某個 User 的 Posts，就要把 userId 擺在 Row key 的前面，如果要按時間排序，那中間要放文章的發布時間，最後因為同一個時間可能會有多篇文章（雖然實際上不太可能，但理論上我們要容許這種情形），所以結尾還要放 postId。

所以 RowKey 可能長這樣 `{userId}-{postReleaseTime}-{postId}`。
如果想要找出某個使用者的文章就下 `{userId}` 開頭的查詢（可能有點像 MySQL 的 `like '{userId}-%'`）。

### NoSQL
#### 問題B

> 設計一個NoSQL DB的rowkey，並說明設計原因，滿足
>   - 在某個latlngbounds時，能快速找出結果
>   - 依照NoSQL DB特性，避免hotspot產生

因為目標的經度與現在位置最接近時，緯度未必也最接近，甚至有可能兩者在都不是離現在位置個別的位置最接近，但實際上距離最近。這邊我找到的解決方式要將經緯度轉換成 GeoHash，這個 Hash 可以幫我們將二維的經緯度轉換成一維來表示，且如果兩個 GeoHash 前綴相符，表示在某個尺度上兩者位於同一個地理空間。

所以 RowKey 可能長這樣 `{geohash}`。
如果想要找出鄰近目前位置的地點，可以用目前的經緯度作 GeoHash，取前面幾位，進資料庫作比對。
實際要取幾位就看地點的密集程度了，這邊我也沒做過（🥲）。


### Coding
#### 程式碼架構說明

##### 抽出 .env
因為 `no other additional third party packages are allowed`，引用環境變數就不能是透過 dotenv，而是啟動時指定。這邊我是寫成 Makefile，透過 Makefile 執行。

##### package main
go 沒有特別限定 `package main` 的檔案擺放在哪裡，社群的慣例好像都是擺在 `cli`、`cmd` 或者`專案根目錄`之類的資料夾。這裡我擺在 `main`。

考慮之後還可能有其他使用情境（entrypoints），譬如 `cli.go`、`messageQueue.go`，事先規劃怎麼將它們分隔開來也是重要的事情。

##### modules & N-Tiers
專案的目錄架構除了垂直的分層，還應該要有水平的、模組化的區分。

垂直分層：
- presentation
    - http
    - messageQueue
- application:
    - commands
    - queries
- domain:
    - models
    - services
- adapters:
    - repositories
    - clients, sessions, services, storages...
- utils: 
    - encrypters / descrypters
    - hashers
    - generators

水平區分：按專案模組化的需求或不同團隊的職責來劃分。

##### ValidateApiKey 為什麼是擺在 `utils/auth` 下
因為目前 `ValidateApiKey` 沒有涉及到資料庫或對第三方 API 的呼叫，倘若之後有更複雜的實作，應該會擺在 `modules/iam/application/queries` 下。