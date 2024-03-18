---
theme: apple-basic
title: YARTBML
info: |
  Yet Another Re-implementation of Thorsten Ball's Monkey Language
highlighter: shiki
drawings:
  persist: false
transition: slide-left
mdc: true
layout: intro
---

# YARTBML
## Yet Another Re-implementation of Thorsten Ball's Monkey Language

<div class="mt-12">
  <span class="font-500">
    By: Dinesh Umasankar, Joesph Porrino, Katherine Banis, Paul Jensen
  </span>
</div>

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/dineshUmasankar/YARTBML" target="_blank" alt="GitHub" title="Open in GitHub"
    class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

---
layout: image-right
transition: slide-left
image: tbook.png
---

# What is YARTBML?

Minimally Functional-Paradigm Inspired Language (Based on SMoL)

<v-clicks>

- 🛠 Built on the foundation provided by Thorsten Ball's: "Writing an Interpreter in Go"
- 💭 Inspired by the many forks of this foundation to provide the best feature set and experience
- 👩‍💻 Focuses on the developer experience for building general-purpose applications
- 🍕 Learnable in a lunch break

</v-clicks>

---
layout: section
---

# **Let's Prove It**

---
---

# File Format

<br/>
<v-clicks>

- To start writing code in YARTBML, make a `.ybml` file.
- All programs will be interpreted under UTF-8, and English Alphanumeric Characters only.
- Let's make an example fibonacci program.

</v-clicks>

---
layout: section
---

# Fibonacci Program

<div class="font-italic text-sm">fibonacci.ybml</div>
```js {1|2-4|5-9|13|all}
let fibonacci = fn(x) {
    if (x == 0) {
        return 0;
    } else {
        if (x == 1) {
            return 1;
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

puts(fibonacci(10)) // Displays "55".
```

---
layout: section
---

# **Core Language Principles**

---
layout: statement
---

# Parsing
Pratt Parsing Technique (form of Recursive Descent)

---
layout: statement
---

# Interpreter / Evaluator
Tree-Walking Interpreter