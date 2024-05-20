# q6

## Preretirement
- golang@1.22.3
- docker
- GNU Make

## Quick start
Start the server
```bash
make build && ./q6
```
And visit http://localhost:8080/swagger/index.html

## Makefile
```bash
# build
make build

# test
make test

# test report
make report

# clean
make clean

# automatically install swag to ./bin and run it
make doc

# automatically install golangci-lint to ./bin and run it
make lint

# build docker image
make image # with default tag name 'latest'
env TAG=1.0.0 make image # or provide a tag name
```

## Project structure
```bash
.
├── bin # develop tools (golangci-lint, swag)
├── docs # swag generating docs
├── lib # libraries for non-business logic
└── pkg # business logic
```

## System design

由於男生只能與較矮的女生進行配對，所以越高的男生越容易配對。同理，越矮的女生也越容易配對。因此，直觀的實現方式是使用 `priority queue` (`heap`)。

系統維護兩個基於身高進行比較的 `heap`：
- 尚未配對男生的 `max heap`
- 尚未配對女生的 `min heap`

系統也支援取消配對請求的操作，所以需要再維護一個 `map`，用來記錄所有的配對請求，以方便檢查其是否存在並從 `heap` 中移除。

### AddSinglePersonAndMatch

先將配對請求放入`map` 中，在依據性別放入其對應的 `heap` 中。
操作 `map` 時間複雜度為 `O(1)`，操作 `heap` 的時間複雜度為 `O(log(n))`，所以 `AddSinglePersonAndMatch` 的時間複雜度為 `O(log(n))`

配對只會在有新的配對請求時發生，因此系統會在執行 `AddSinglePersonAndMatch` 後進行配對。

配對邏輯如下：只要 "最高男生的身高" 大於 "最矮女生的身高"，即可配對。需要注意的是，一個人可以進行多次配對，因此需要根據配對請求的數量，取出多個來檢查是否符合配對條件。

### RemoveSinglePerson

`RemoveSinglePerson` 使用 `map` 找到特定用戶，並將其從對應的 `heap` 中移除。其操作的時間複雜度為 `O(log(n))`。

### QuerySinglePeople

此方法根據請求的數量 `N`，從對應的 `heap` 中取出並查看，再放回去即可。其操作的時間複雜度為 `O(N * log(n))`。

### 各項方法的時間複雜度總結
| 方法                      | 時間複雜度        |
| ------------------------- | ---------------- |
| AddSinglePersonAndMatch   | O(log(n))        |
| RemoveSinglePerson        | O(log(n))        |
| QuerySinglePeople         | O(N * log(n))    |

(`n` 為該性別尚未配對的人數; `N` 為請求的數量)
