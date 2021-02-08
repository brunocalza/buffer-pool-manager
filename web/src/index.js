import {Svg2Roughjs, RenderMode} from 'svg2roughjs'
import * as d3 from 'd3'
import {graphviz} from "d3-graphviz";

const svg2roughjsDisk = new Svg2Roughjs('#disk-rough', RenderMode.SVG, {
    seed : 1000
})

const svg2roughjsBP = new Svg2Roughjs('#buffer-pool-rough', RenderMode.SVG, {
    seed : 1000
})


const svg2roughjsCR = new Svg2Roughjs('#clock-replacer-rough', RenderMode.SVG, {
    seed : 1000
})

function renderBufferPool(pagesTable, maxPoolSize, pinCount) {
    let dotString = `
    digraph BufferPool {
         Description [shape="plaintext", label="Buffer Pool"]

          PageTable[
           shape = "plaintext"
            label=<
           <table border='0' cellborder="1" >
           ${Object.keys(pagesTable).map(page => `<tr><td ${pinCount[page] > 0 ? 'color="red"' : ''} colspan="1" port='${page}' > Page ${page} </td></tr>`).join("")}
           </table>
           >];
           
          Pages[
           shape = "plaintext"
            label=<
           <table border='0' cellborder="1" >
                ${[...Array(maxPoolSize).keys()].map(frame => `<tr><td colspan="1" port='${frame}' > Frame ${frame} </td></tr>`).join("")}
           </table>
           >];
           
           ${Object.keys(pagesTable).map(page => `PageTable:${page} -> Pages:${pagesTable[page]}`).join("")}
          
          {rank = same; PageTable; Pages;}
      }
    `
    
    graphviz("#buffer-pool")
    .renderDot(dotString, function(){
        const svg = document.querySelector('#buffer-pool svg')
        svg2roughjsBP.svg = svg
        svg2roughjsBP.fontFamily = 'Indie Flower'
        svg2roughjsBP.randomize = false
        svg2roughjsBP.pencilFilter = true
        svg2roughjsBP.redraw()

        const outputSvg = document.querySelector('#buffer-pool-rough svg')
        outputSvg.removeChild(outputSvg.childNodes[1]) 
    });
}

function renderDisk(pagesInDisk, maxDiskNumPages) {
    let dotString = `digraph Disk {  
        Description [shape="plaintext", label="Disk"]
            
            Pages[
            shape = "plaintext" 
            label=<
            <table border='0' cellborder="1"  >
                <tr >
                ${[...Array(maxDiskNumPages).keys()].map(page => `<td ${pagesInDisk.includes(page) ? '' : 'style="invis"'} colspan="1" port='${page}' > Page ${page} </td>`).join("")}
                </tr>
            </table>
            >];
        }`

    console.log(dotString)

    graphviz("#disk")
    .renderDot(dotString, function(){
        const svg = document.querySelector('#disk svg')
        svg2roughjsDisk.svg = svg
        svg2roughjsDisk.fontFamily = 'Indie Flower'
        svg2roughjsDisk.randomize = false
        svg2roughjsDisk.pencilFilter = true
        svg2roughjsDisk.redraw()

        const outputSvg = document.querySelector('#disk-rough svg')
        outputSvg.removeChild(outputSvg.childNodes[1]) 
    });
}

function renderClockReplacer(clockReplacer) {
    let dotString = `
    digraph ClockReplacer {

        Description [shape="plaintext", label="Clock Replacer", pos="0,1.5!"]

        Pointer [label="CP", shape="plaintext", pos="0,0!"]

    0[
        ${clockReplacer['Clock'][0] === undefined ? 'style="invis"' : ''}
        pos="6.123233995736766e-17,1.0!", shape = "plaintext"
        label=<
        <table border='0' >
        <tr>
                <td colspan="1" port='ref' > ref = ${clockReplacer['Clock'][0] !== undefined ? clockReplacer['Clock'][0]["ReferenceValue"] ? "1" : "0" : ""} </td>
        </tr>
        <tr>
                <td border='1' colspan="1" port='frame'>${clockReplacer['Clock'][0] !== undefined ? clockReplacer['Clock'][0]["ClockFrame"] : ""} </td>

            </tr>
        </table>
        >];
        
    3[
        ${clockReplacer['Clock'][3] === undefined ? 'style="invis"' : ''}
        pos="-1.0,1.2246467991473532e-16!", shape = "plaintext"
        label=<
        <table border='0' >
        <tr>
                <td colspan="1" port='ref' > ref = ${clockReplacer['Clock'][3] !== undefined ? clockReplacer['Clock'][3]["ReferenceValue"] ? "1" : "0" : ""} </td>
        </tr>
        <tr>
                <td border='1' colspan="1" port='frame'>${clockReplacer['Clock'][3] !== undefined ? clockReplacer['Clock'][3]["ClockFrame"] : ""} </td>

            </tr>
        </table>
        >];
        
    2[
        ${clockReplacer['Clock'][2] === undefined ? 'style="invis"' : ''}
        pos="-1.8369701987210297e-16,-1.0!", shape = "plaintext"
        label=<
        <table border='0' >
        <tr>
                <td colspan="1" port='ref' > ref = ${clockReplacer['Clock'][2] !== undefined ? clockReplacer['Clock'][2]["ReferenceValue"] ? "1" : "0" : ""} </td>
        </tr>
        <tr>
                <td border='1' colspan="1" port='frame'>${clockReplacer['Clock'][2] !== undefined ? clockReplacer['Clock'][2]["ClockFrame"] : ""} </td>

            </tr>
        </table>
        >];
        
    1[
        ${clockReplacer['Clock'][1] === undefined ? 'style="invis"' : ''}
        pos="1.0,-2.4492935982947064e-16!", shape = "plaintext"
        label=<
        <table border='0' >
        <tr>
                <td colspan="1" port='ref' > ref = ${clockReplacer['Clock'][1] !== undefined ? clockReplacer['Clock'][1]["ReferenceValue"] ? "1" : "0" : ""} </td>
        </tr>
        <tr>
                <td border='1' colspan="1" port='frame'>${clockReplacer['Clock'][1] !== undefined ? clockReplacer['Clock'][1]["ClockFrame"] : ""} </td>

            </tr>
        </table>
        >];

    Pointer -> ${clockReplacer['ClockHand']}:frame:s
    }`

    console.log(dotString)

    graphviz("#clock-replacer")
            .engine("neato")
            .renderDot(dotString, () => {  
                const svg = document.querySelector('#clock-replacer svg')
                svg2roughjsCR.svg = svg
                svg2roughjsCR.fontFamily = 'Indie Flower'
                svg2roughjsCR.randomize = false
                svg2roughjsCR.pencilFilter = true
                svg2roughjsCR.redraw()

                const outputSvg = document.querySelector('#clock-replacer-rough svg')
                outputSvg.removeChild(outputSvg.childNodes[1])
            })
}

document.getElementById("new-page").onclick = function() {
    fetch('http://localhost:3000/new')
    .then(response => response.json())
    .then(data => { 
        renderDisk(data['PagesInDisk'], data['MaxDiskNumPages'])
        renderBufferPool(data['PagesTable'], data['MaxPoolSize'], data['PinCount'])
        renderClockReplacer(data['ClockReplacer'])
    })
    .catch((err) => {
        console.log(err)
    });
}

document.getElementById("flush-page").onclick = function() {
    var pageId = prompt("Enter the page id");

    if (pageId != null) {
        fetch(`http://localhost:3000/flush?page=${pageId}`)
        .then(response => response.json())
        .then(data => { 
        renderDisk(data['PagesInDisk'], data['MaxDiskNumPages'])
        renderBufferPool(data['PagesTable'], data['MaxPoolSize'], data['PinCount'])
        renderClockReplacer(data['ClockReplacer'])
        })
        .catch((err) => {
            console.log(err)
        });
    }
}

document.getElementById("delete-page").onclick = function() {
    var pageId = prompt("Enter the page id");

    if (pageId != null) {
        fetch(`http://localhost:3000/delete?page=${pageId}`)
        .then(response => response.json())
        .then(data => { 
        renderDisk(data['PagesInDisk'], data['MaxDiskNumPages'])
        renderBufferPool(data['PagesTable'], data['MaxPoolSize'], data['PinCount'])
        renderClockReplacer(data['ClockReplacer'])
        })
        .catch((err) => {
            console.log(err)
        });
    }
}

document.getElementById("unpin-page").onclick = function() {
    var pageId = prompt("Enter the page id");

    if (pageId != null) {
        fetch(`http://localhost:3000/unpin?page=${pageId}`)
        .then(response => response.json())
        .then(data => { 
        renderDisk(data['PagesInDisk'], data['MaxDiskNumPages'])
        renderBufferPool(data['PagesTable'], data['MaxPoolSize'], data['PinCount'])
        renderClockReplacer(data['ClockReplacer'])
        })
        .catch((err) => {
            console.log(err)
        });
    }
}

document.getElementById("fetch-page").onclick = function() {
    var pageId = prompt("Enter the page id");

    if (pageId != null) {
        fetch(`http://localhost:3000/fetch?page=${pageId}`)
        .then(response => response.json())
        .then(data => { 
        renderDisk(data['PagesInDisk'], data['MaxDiskNumPages'])
        renderBufferPool(data['PagesTable'], data['MaxPoolSize'], data['PinCount'])
        renderClockReplacer(data['ClockReplacer'])
        })
        .catch((err) => {
            console.log(err)
        });
    }
}

document.getElementById("flush-all").onclick = function() {
    fetch(`http://localhost:3000/flush-all`)
    .then(response => response.json())
    .then(data => { 
        renderDisk(data['PagesInDisk'], data['MaxDiskNumPages'])
        renderBufferPool(data['PagesTable'], data['MaxPoolSize'], data['PinCount'])
        renderClockReplacer(data['ClockReplacer'])
    })
    .catch((err) => {
        console.log(err)
    });
}



let dotString = `
digraph ClockReplacer {

    Description [shape="plaintext", label="Clock Replacer", pos="0,1.5!"]

    Pointer [label="CP", shape="plaintext", pos="0,0!"]

  0[
      style="invis"
      pos="6.123233995736766e-17,1.0!", shape = "plaintext"
      label=<
     <table border='0' >
       <tr>
            <td colspan="1" port='ref' > ref = 1 </td>
      </tr>
       <tr>
            <td border='1' colspan="1" port='frame'>0</td>

        </tr>
     </table>
     >];
     
  3[
      style="invis"
      pos="-1.0,1.2246467991473532e-16!", shape = "plaintext"
      label=<
     <table border='0' >
       <tr>
            <td colspan="1" port='ref' > ref = 1 </td>
      </tr>
       <tr>
            <td border='1' colspan="1" port='frame'>3</td>

        </tr>
     </table>
     >];
     
  2[
      style="invis"
      pos="-1.8369701987210297e-16,-1.0!", shape = "plaintext"
      label=<
     <table border='0' >
       <tr>
            <td border='1' colspan="1" port='frame'>2</td>

        </tr>
        <tr>
            <td colspan="1" port='ref' > ref = 1 </td>
        </tr>
     </table>
     >];
     
  1[
      style="invis"
      pos="1.0,-2.4492935982947064e-16!", shape = "plaintext"
      label=<
     <table border='0' >
       <tr>
            <td colspan="1" port='ref' > ref = 1 </td>
      </tr>
       <tr>
            <td border='1' colspan="1" port='frame'>1</td>

        </tr>
     </table>
     >];

  Pointer -> 0:frame:s
}`
graphviz("#clock-replacer")
        .engine("neato")
        .renderDot(dotString, () => {  
            const svg = document.querySelector('#clock-replacer svg')
            svg2roughjsCR.svg = svg
            svg2roughjsCR.fontFamily = 'Indie Flower'
            svg2roughjsCR.randomize = false
            svg2roughjsCR.pencilFilter = true
            svg2roughjsCR.redraw()

            const outputSvg = document.querySelector('#clock-replacer-rough svg')
            outputSvg.removeChild(outputSvg.childNodes[1])
        })



