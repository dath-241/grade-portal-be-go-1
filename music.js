const audio = document.getElementById('audio');

audio.play().catch(error => {
  console.error('Không thể phát nhạc:', error);
});