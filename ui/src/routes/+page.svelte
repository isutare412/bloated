<script lang="ts">
	import jerryThanks from '$assets/jerryThanks.gif'
	import CaptionImage from '$components/image/CaptionImage.svelte'
</script>

<article class="prose prose-code:bg-secondary prose-code:text-secondary-content w-full max-w-none">
	<h1>Yet Another Blog</h1>
	<CaptionImage src={jerryThanks} alt="Jerry bowing at a stage" caption="Hello, World!" />
	<p>
		Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
		labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
		laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in
		voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
		non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	</p>
	<hr />
	<p>We need to use <code>http.Client</code>.</p>
</article>
