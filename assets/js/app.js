var axios = require('axios');

var getTodos = () => {
  return axios.get('/api/v1/todos')
}

(function(Vue) {
  var todoComponent = Vue.component('todo', {
    props: ['todo'],
    template: `
        <label>
            <input type="checkbox" v-on:click="checkedTodo(todo)"> {{ todo.text }}
        </label>
        `,
    methods: {
      checkedTodo: function(todo) {
        axios.put(`/api/v1/checked/${todo.todo_id}`).then((res) => {
          todo.status = true;
        });
      }
    }
  })

  var todoItemsComponent = Vue.component('todo-items', {
    props: ['todoStatus', 'todos'],
    template: `
      <ul>
          <div v-for="todo in todos">
              <li v-if="todo.status==todoStatus">
                  <todo :todo="todo"></todo>
              </li>
          </div>
      </ul>
    `,
    components: {
      todoComponent,
    },
  })

  var todoInputComponent = Vue.component('todo-input-field', {
    props: ['todos'],
    template: `
      <div>
          <textarea v-model="content"></textarea>
          <button @click="addTodo">add Todo</button>
      </div>
      `,
    data: function() {
      return {
        content: ''
      }
    },
    methods: {
      addTodo: function() {
        axios.post("/api/v1/todo", {
          text: this.content
        }).then((res) => {
            this.$emit('reload');
        })
        this.content = "";
      }
    }
  })

  var todoAppComponent = Vue.component('todo-app', {
    template: `
      <div>
          <todo-input-field :todos="todos" v-on:reload="reloadTodoList"></todo-input-field>
          <h3>TODO</h3>
          <todo-items :todo-status=false :todos="todos"></todo-items>
          <h3>DONE</h3>
          <todo-items :todo-status=true :todos="todos"></todo-items>
      </div>
      `,
    components: {
      todoInputComponent,
    },
    created: function() {
      getTodos().then((res) => {
        this.todos = res.data;
      }).catch((res) => {
        console.log("error");
      })
    },
    data: function() {
      return {
        todos: []
      }
    },
    methods: {
        reloadTodoList: function() {
            console.log("reloadTodoList")
            getTodos().then((res) => {
                this.todos = res.data;
            }).catch((res) => {
                console.log("reloadTodoList error");
            })
        }
    }
  })

  // TODO: refactaring
  new Vue({
    el: '#app',
    components: {
      todoAppComponent,
    }
  });
})(Vue);
