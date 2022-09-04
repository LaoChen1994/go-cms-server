import Path from 'path'
import { Configuration } from 'webpack'
import Glob from 'glob'
import Util from 'util'

export default async function getResolves(): Promise<Configuration['resolve']> {
  const files = await Util.promisify(Glob)("src/**")

  const alias = files.reduce((p: Record<string, string>, name: string) => {
    const dirname = name.match(/src\/(.*)/)
    if (!dirname) return p;

    p[`${dirname[1][0].toUpperCase()}${dirname[1].slice(1)}`] = Path.resolve(__dirname, "../../", name)

    return p
  }, {} as Record<string, string>)

  return {
    alias,
    extensions: [".tsx", ".ts", ".js", ".scss", ".sass", ".css"],
  }
}
