## Skills
A skill is a local instruction set stored in a `SKILL.md` file. Use skills to make frontend tasks in this repository consistent and repeatable.

### Available skills
- `vue-view-builder`: Build or refactor Vue 3 pages (`<script setup>`) and wire routes in this repo. Use when adding new views, restructuring existing pages, or implementing page-level interactions. (file: `.agents/skills/vue-view-builder/SKILL.md`)
- `api-service-builder`: Add or update frontend API modules based on `src/utils/request.js` and integrate them into views/stores. Use when implementing data fetching, submit flows, or error-handling behavior. (file: `.agents/skills/api-service-builder/SKILL.md`)
- `pinia-store-pattern`: Design or refactor Pinia stores (including auth/token persistence) with the current project conventions. Use when adding shared state, auth state, or store actions. (file: `.agents/skills/pinia-store-pattern/SKILL.md`)
- `element-plus-admin-ui`: Implement admin-style forms/tables/dialogs with Element Plus and project-friendly SCSS. Use when building CRUD UI, search/filter forms, list pages, or edit dialogs. (file: `.agents/skills/element-plus-admin-ui/SKILL.md`)

### How to use skills
- Discovery: Read this list first, then open only the required `SKILL.md` files.
- Trigger rules: If a user names a skill directly or the request clearly matches a skill description, use that skill for the turn.
- Multiple skills: Use the minimal combination needed for the request, and state the execution order briefly.
- Progressive loading: Load only referenced files that are needed for the task; avoid bulk-reading unrelated files.
- Fallback: If a skill file is missing or unclear, state the issue briefly and continue with the closest project convention.

## Project baseline
- Stack: Vue 3 + Vite + Element Plus + Pinia + Vue Router + Axios.
- Entry points: `src/main.js`, `src/router/index.js`, `src/stores/index.js`, `src/utils/request.js`.
- Alias: Use `@` to reference `src`.
- Styling: Prefer local `scoped` SCSS for page/component styles and keep styles close to the component.
- API client: Reuse `src/utils/request.js` interceptors and avoid duplicating global auth/error handling.
