// Пример данных о пациентах
const patients = [
    {
        ID: 1,
        FirstName: "Иван",
        LastName: "Иванов",
        PlansTherapy: [
            { ID: 1, StartDate: "2024-10-01", FinishDate: "2024-10-05", Description: "Химиотерапия", SideEffect: "Усталость" },
            { ID: 2, StartDate: "2024-10-10", FinishDate: "2024-10-15", Description: "Лучевая терапия", SideEffect: "Тошнота" },
        ],
    },
    {
        ID: 2,
        FirstName: "Мария",
        LastName: "Петрова",
        PlansTherapy: [
            { ID: 1, StartDate: "2024-10-05", FinishDate: "2024-10-10", Description: "Хирургическое вмешательство", SideEffect: "Боль" },
        ],
    }
];

let currentYear = new Date().getFullYear();
let currentMonth = new Date().getMonth();

document.addEventListener('DOMContentLoaded', function() {
    generateCalendar(currentYear, currentMonth, patients);
    updateCurrentMonth();

    document.getElementById('prevMonth').onclick = () => {
        currentMonth--;
        if (currentMonth < 0) {
            currentMonth = 11;
            currentYear--;
        }
        generateCalendar(currentYear, currentMonth, patients);
        updateCurrentMonth();
    };

    document.getElementById('nextMonth').onclick = () => {
        currentMonth++;
        if (currentMonth > 11) {
            currentMonth = 0;
            currentYear++;
        }
        generateCalendar(currentYear, currentMonth, patients);
        updateCurrentMonth();
    };
});

function updateCurrentMonth() {
    const monthNames = ["Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"];
    document.getElementById('currentMonth').textContent = `${monthNames[currentMonth]} ${currentYear}`;
}

function generateCalendar(year, month, patients) {
    const calendarContainer = document.getElementById('calendar');
    calendarContainer.innerHTML = '';

    const firstDay = new Date(year, month, 1);
    const lastDay = new Date(year, month + 1, 0);
    const daysInMonth = lastDay.getDate();
    const today = new Date();

    const startOffset = (firstDay.getDay() + 6) % 7; // Смещение, чтобы понедельник был первым

    // Дни недели
    const weekDays = ["Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"];
    weekDays.forEach(day => {
        const dayLabel = document.createElement('div');
        dayLabel.className = 'cell day-label';
        dayLabel.textContent = day;
        calendarContainer.appendChild(dayLabel);
    });

    for (let i = 0; i < startOffset; i++) {
        const emptyCell = document.createElement('div');
        emptyCell.className = 'cell';
        calendarContainer.appendChild(emptyCell);
    }

    for (let day = 1; day <= daysInMonth; day++) {
        const dayElement = document.createElement('div');
        dayElement.className = 'cell';
        dayElement.textContent = day;

        if (day === today.getDate() && month === today.getMonth() && year === today.getFullYear()) {
            dayElement.classList.add('current-day'); // Подсветка текущего дня
        }

        patients.forEach((patient, index) => {
            const plansForDay = getPlansForDay(day, month + 1, year, [patient]);
            plansForDay.forEach(plan => {
                const eventElement = document.createElement('div');
                eventElement.className = `event patient-${index + 1}`; // Добавление класса для пациента
                eventElement.textContent = `${patient.LastName}: ${plan.Description}`; // Фамилия пациента + текст события
                dayElement.appendChild(eventElement);
            });
        });

        calendarContainer.appendChild(dayElement);
    }
}

function getPlansForDay(day, month, year, patients) {
    const plans = [];
    patients.forEach(patient => {
        patient.PlansTherapy.forEach(plan => {
            const startDate = new Date(plan.StartDate);
            const finishDate = new Date(plan.FinishDate);
            if (startDate.getDate() <= day && finishDate.getDate() >= day &&
                startDate.getMonth() + 1 === month && startDate.getFullYear() === year) {
                plans.push(plan);
            }
        });
    });
    return plans;
}

document.getElementById('closeModal').onclick = () => {
    document.getElementById('modal').style.display = 'none';
};

window.onclick = (event) => {
    if (event.target === document.getElementById('modal')) {
        document.getElementById('modal').style.display = 'none';
    }
};
