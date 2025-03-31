import { defineCollection, z } from "astro:content";
import { glob } from "astro/loaders";

const tags = z.enum(["store"]);

const posts = defineCollection({
	loader: glob({ pattern: "**/*.md", base: "./src/posts" }),
	schema: z.object({
		title: z.string(),
		date: z.date(),
		tags: z.array(tags),
	}),
});

export const collections = { posts };
