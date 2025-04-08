import { createWindowSize } from "@solid-primitives/resize-observer";
import type { InferEntrySchema, RenderedContent } from "astro:content";
import { createSignal } from "solid-js";

const sm = 640;
const md = 768;
const lg = 1024;
const xl = 1280;
const twoxl = 1536;

const Nav = ({ posts, path }: {
    posts: {
        id: string;
        body?: string;
        collection: "posts";
        data: InferEntrySchema<"posts">;
        rendered?: RenderedContent;
        filePath?: string;
    }[], path: string
}) => {
    const windowSize = createWindowSize();
    const [exp, setExp] = createSignal(false);

    return windowSize.width <= lg ? <div class={`w-full bg-bg z-40 border-fg border-b-4 top-0 absolute ${exp() ? "max-h-300 h-full" : "h-24"} transition-all duration-500 ease-in-out`}>
        <div class="p-6 flex flex-row justify-between">
            <a href="/" class="flex flex-col space-y-2 justify-center w-full">
                <span class="text-4xl"><span class="text-orange">⦿</span>¬<span class="text-cyan">♜</span></span>
            </a>
            <button class={`text-5xl ${exp() ? "rotate-45" : ""} trasition-all duration-100`} onClick={() => setExp(!exp())}>
                +
            </button>
        </div>
        <div class={` border-fg flex flex-col space-y-8 ${exp() ? "opacity-100" : "opacity-0"} transition-all duration-500 ease-in-out`}>
            <ol class="w-full h-full flex flex-col border-t-4 border-fg">
                {
                    posts.map(post => {
                        return (
                            <li class={
                                `border-b-4 border-fg ${`/${post.id}` === path ? "bg-orange text-bg" : "hover:bg-fg hover:text-bg"} flex`
                            }>
                                <a class="font-serif font-bold text-xl p-4 w-full h-full" href={`/${post.id}`}>{post.data.title}</a>
                            </li>
                        )
                    })}
            </ol>
        </div>
    </div> :
        <div class={`border-r-8 border-fg md:w-1/6 flex flex-col space-y-8`}>
            <div class="flex flex-col space-y-8 p-8">
                <a href="/" class="flex flex-col space-y-2 justify-center w-full">
                    <span class="text-5xl"><span class="text-orange">⦿</span>¬<span class="text-cyan">♜</span></span>
                    <span class="font-mono font-bold">
                        (<span class="text-orange">checkers</span> not <span class="text-cyan">chess</span>)
                    </span>
                </a>
                <div class="flex flex-col space-y-2">
                    <p class="font-sans border-l-4 border-fg px-2">Solving problems you probably don't have creates more problems you definitely do.</p>
                    <p class="font-serif font-bold">- Mike Acton</p>
                </div>
            </div>
            <ol class="w-full h-full flex flex-col border-t-4 border-fg">
                {
                    posts.map(post => {
                        console.log(post.id);
                        return (
                            <li class={
                                `border-b-4 border-fg ${`/${post.id}` === path ? "bg-orange text-bg" : "hover:bg-fg hover:text-bg"} flex`
                            }>
                                <a class="font-serif font-bold text-xl p-4 w-full h-full" href={`/${post.id}`}>{post.data.title}</a>
                            </li>
                        )
                    })}
            </ol>
        </div>
}
export default Nav;