$(() => {
    if (!isNotFound()) {
        let bucket = findCurrentBucket();
        $.get(`/aws-s3?b=${bucket}`)
            .done(data => {
                if (data.list) {
                    setImageCount(data.list.length);
                    data.list.forEach(function (_, i) {
                        appendImage(data, i);                 
                        setTimeout(function () {
                            displayImages(data)
                        }, 3000);
                    });
                } else {
                    setImageCount(0); 
                }
            })
    } else {
        $('footer').html('');
    }
})

function findCurrentBucket() {
    let parser = document.createElement('a');
    parser.href = window.location.href
    let bucket = `webcam-${parser.pathname.substring(1)}`;
    if (bucket === 'webcam-') {
        bucket += 'timelapse'
    }
    return bucket
}

function setImageCount(count) {
    $('footer').html(`[${count}] images`)
}

function isNotFound() {
    return $('header').html() != "";
}

function displayImages(data) {
    let max = data.list.length;
    let current = 0;
    let loading = true
    setInterval(function () {
        if (current >= max) {
            current = 0;
        }
        let img = document.querySelector(`#image-${current}`);
        if (img.complete) {
            if (loading) {
                $('.loading').hide();
                loading = false
            }
            displayNext(img, current, max)   
            current++
        }
    }, 200);
}

function displayNext(next, index, max) {
    let previous = index - 1;
    if (previous < 0) {
        previous = max - 1
    }
    let img = document.querySelector(`#image-${previous}`);
    next.classList.remove('next');;
    next.classList.add('front');
    img.classList.add('next');
    img.classList.remove('front');
}

function appendImage(images, index) {
    let url = `${images['base-url']}/${images.name}/${images.list[index].Key}`;
    $('main').append(`<img class="next" id="image-${index}" src="${url}">`) 
}
