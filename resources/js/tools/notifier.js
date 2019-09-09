import Swal from 'sweetalert2'

class Notifier {

    constructor() {

        this.toast = Swal.mixin({
            toast: true,
            position: 'top-end',
            showConfirmButton: false,
            timer: 3000
        });     
    }

    success(text) {
        this.toast.fire({
            title: text,
            type: 'success',
        });
    }

    info(text) {
        this.toast.fire({
            title: text,
            type: 'info',
        });
    }

    warning(text) {
        this.toast.fire({
            title: text,
            type: 'warning',
        });
    }

    error(text) {
        this.toast.fire({
            title: text,
            type: 'error',
        });
    }
}

const notifier = new Notifier()

export default notifier