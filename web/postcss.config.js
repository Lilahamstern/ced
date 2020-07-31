const postcssimport = require("postcss-import");
const postcsspresetenv = require("postcss-preset-env");
const postcssnested = require("postcss-nested");
const autoprefixer = require("autoprefixer");
const tailwindcss = require("tailwindcss");
const postcssPurgecss = require("@fullhuman/postcss-purgecss");

const purgecss = postcssPurgecss({
  content: ["./public/**/*.html", "./src/**/*.tsx", "./src/**/*.css"],
  defaultExtractor: (content) => content.match(/[\w-/:]+(?<!:)/g) || [],
  whitelistPatterns: [
    /-(leave|enter|appear)(|-(to|from|active))$/,
    /^(?!(|.*?:)cursor-move).+-move$/,
    /^router-link(|-exact)-active$/,
  ],
});

module.exports = {
  plugins: [
    postcssimport,
    tailwindcss,
    postcsspresetenv({
      stage: 1,
    }),
    postcssnested,
    autoprefixer,
    ...(process.env.NODE_ENV === "production" ? [purgecss] : []),
  ],
};
