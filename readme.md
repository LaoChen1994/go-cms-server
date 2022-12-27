# GO CMS Server

## 简介

一个前端新手小白学习GO Web的练手项目，主要技术栈

前端：React + Mobx

后端：Go Gin + Go Orm

## MD Compile

### 1. 预备知识

#### 1.1 需要用到的库

1. remark
2. rehype
3. unified

如果我们在研究过MDX或MDsveX，我们会发现其底层用的都是`unified`生态中的`remark`和`rehype`

因此我们只需要使用`remark`或者给`rehype`处理器就可以实现我们自己的`MDX`转换器

#### 1.2 什么是remark 或 rehype

[官网地址](https://unifiedjs.com/explore/package/unified/)

其实我们可以简单理解，其主要就是 `AST解析语法树` + `插件解析语法进行渲染`

从中我们可以直到其工作主要分为两部分：

1. 使用预处理器解析: `string` => `ast`

   ```markdown
   # 你好 
   ```
   经过ast转换后，其结果如下：
   ```json
   {
      "type": "heading",
      "depth": 1,
      "children": [
         {
            "type": "text",
            "value": "你好"       
         }
      ]
   }
   ```
2. 使用插件进行渲染: `ast` => `react component` / `html`，上述的ast转换后得到结果如下

   ```html
   <h1>你好</h1>
   ```

### 2. Remark/Rehype如何处理Markdown解析的上述步骤

1. 使用`unified`处理器
2. 使用`remark-parse` / ``处理得到markdown的ast语法树
3. 得到`mdast`语法树
4. 使用`插件`转换语法树
5. 使用`remark-stringify` / `rehype-stringify`序列化语法树到字符串

### 3. 何时使用Remark / Rehype

1. 如果处理更偏向于markdown语法树，使用`remark`
2. 如果插件处理的是最后输出的`html`使用`rehype`
