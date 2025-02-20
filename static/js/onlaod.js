
function main() {
    const cards = document.querySelectorAll(".artist-card div")
    cards.forEach(card => {
        const img = card.querySelector("img")
        // img.onload = () => card.style.width = "auto"; card.style.height = "auto"; card.style.padding = "20px"
    });
}

main()