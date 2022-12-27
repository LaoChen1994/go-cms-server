import { unified } from 'unified';
import remarkParse from "remark-parse";
import remarkFrontmatter from "remark-frontmatter";
import remarkGfm from "remark-gfm";
import remarkRehype from "remark-rehype";
import rehypeStringify from "rehype-stringify";

import { plugin } from './plugin'

async function MDCompile(mdString: string): Promise<string> {
  const processor = unified()
    .use(
      remarkParse,
    )
    .use(plugin)
    .use(remarkFrontmatter)
    .use(remarkGfm)
    .use(remarkRehype)
    .use(rehypeStringify)

  const result = await processor.process(mdString)

  return result.toString()
}

MDCompile("#123 ## 123 <Hello />")

export default MDCompile
