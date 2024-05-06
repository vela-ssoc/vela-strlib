# vela-strlib
基础字符串处理库文件

## 内置方法

- strlib.utf8
- strlib.similarity
- strlib.ac
- strlib.gen

## 用法样例

```lua
    local strlib = vela.strlib             -- 导入方法


    local v1 = "你ab好,cc创"
    local v2 = "好你,创新"

    local u1 = strlib.utf8(v1)             -- utf8 编码 
    local u2 = strlib.utf8(v2)             -- utf8 编码

    u1.trim(19968 , 40959)                 -- trim(起始 , 结束 , 取反) 代码 字符数组范围
    u2.trim(19968 , 40959)                 -- trim(起始 , 结束 , 取反) 代码 字符数组范围


    --u1.trim(97 , 112 , true)
    --u2.trim(97 , 112 , true)

    vela.Debug("%s" , u1.text)           -- u.text 获取字符串 trim 结果
    vela.Debug("%s" , u2.text)           -- u.text 获取字符串


    local similarity = strlib.similarity   -- 字符串相似匹配
    if similarity(u1.text , u2.text).prop > 0.6 then
        vela.Debug("hit > 0.6")
    end

    local black = strlib.ac({"你好", "很好", "还好"} , true) -- true: 是否返回命中结果
    local ok , ret = black("今天的天气还好,心情很好!")
    vela.Debug("%v , %s" , ok , ret)
    
```

## 相似度算法
- levenshtein
- hamming
- jaro
- jacc
- jaro_winkler
- overlap
- smithwatermangotoh
- sorensendice
- prop

```lua
    local strlib = vels.strlib
    local prop = strlib.similarity("a" , "b").prop
    local leve = strlib.similarity("a" , "b").levenshtein
    local hamm = strlib.similarity("a" , "b").hamming
    local jaro = strlib.similarity("a" , "b").jaro
    -- 类似调用 
```

## 字符泛化
支持字符串泛化 结合 相似度匹配 
```lua
    local strlib = vela.strlib
    local gen = strlib.gen
    local emc = vela.attach("员工.txt")

    local new_str = gen("1.explorer.exe>1022.chrome.exe>1023.chrome.exe")
                    .num()                   -- 替换数据为N
                    .graphic(true)           -- 只保留可见字符  , false 取反
                    .space()                 -- 替换空格
                    .file(emc.file , "D")    -- 存在字典内容 替换为D
                    .ip("IP")                -- 替换IP地址为 IP
                    .regex("vela" , "x")     -- 正则替换
                    .gen()                   -- 换成结果
    

    print(new_str)
```