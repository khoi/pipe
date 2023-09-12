import './style.css'
import App from './App.svelte'
import { ListManifests } from '../wailsjs/go/main/App'

const app = new App({
  target: document.getElementById('app')
})

export default app

window.ListManifests = ListManifests
