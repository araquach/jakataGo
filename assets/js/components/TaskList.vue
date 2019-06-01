<template>
    <div>
        <h1>Task List</h1>
        <task-ind v-for="(index, IndTask) in TaskList">
            {{}}
        </task-ind>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                IndTasks: ["Buy milk", "Open Post", "Walk Dog"]
            }
        }
    }
</script>