import { visit } from 'unist-util-visit'
import { Data, Node } from 'unist'

export function transformer(ast: Node<Data>) {
  console.log(ast)
  visit(ast, 'html', (node: Node<any>) => {
    console.log("html =>", node)
  })

  visit(ast, 'link', (node: Node<any>) => {
    console.log('link', node)
  })
}

export function plugin() {
  return transformer
}
