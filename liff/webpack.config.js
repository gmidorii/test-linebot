module.exports = {
  entry: "./src/index.tsx",
  output: {
    filename: "bundle.js",
    path: __dirname + "/public"
  },
  devtool: "source-map",
  mode: "development",
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json", ".css"]
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: [
         'style-loader',
          {
            loader: 'css-loader',
            options: {
              url: false,
              modules: true,
              sourceMap: true,
            }
          },
        ]
      },
      { test: /\.tsx?$/, loader: "awesome-typescript-loader"},
      { enforce: "pre", test: /\.js$/, loader: "source-map-loader"},
    ]
  },
  externals: {
    "react": "React",
    "react-dom": "ReactDOM"
  }
}