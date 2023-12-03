-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
UPDATE academic_programmes SET 
description = 'Profesionales en administración organizan recursos y toman decisiones para que las empresas funcionen de manera eficiente. Supervisan equipos, gestionan finanzas y estrategias. Su impacto radica en el crecimiento y éxito de las empresas, contribuyendo al desarrollo económico y generando empleo.' WHERE id = 1;

UPDATE academic_programmes SET 
description = 'Enfocada en optimizar operaciones industriales, la carrera aborda la gestión eficiente de recursos, mejora de procesos y aseguramiento de calidad en entornos de producción. Contribuye al desarrollo económico al potenciar la eficiencia y competitividad de las empresas manufactureras.'
WHERE id = 2;

UPDATE academic_programmes SET
description = 'Biología estudia seres vivos y procesos. Biólogos exploran células, ecosistemas, contribuyendo a avances médicos, conservación y comprensión de la diversidad biológica, impactando positivamente en salud y sostenibilidad.'
WHERE id = 3;

UPDATE academic_programmes SET 
description = 'Centrada en la transmisión eficaz de información, esta carrera resalta en la creación de contenido relevante. Su impacto es clave para la formación de opiniones, construcción de relaciones y conexión cultural en la sociedad.'
WHERE id = 4;

UPDATE academic_programmes SET
description = 'En el ámbito financiero, los contadores públicos registran y analizan datos económicos. Su labor asegura la transparencia y cumplimiento de normativas contables en empresas. Contribuyen a decisiones financieras sólidas, facilitando el crecimiento económico y la estabilidad empresarial.'
WHERE id = 5;

UPDATE academic_programmes SET
description = 'La carrera de Derecho se enfoca en comprender y aplicar leyes. Los abogados defienden y representan a individuos y empresas. Su impacto es vital para la justicia, resolución de conflictos y mantenimiento del orden social, contribuyendo a sociedades más justas y equitativas.'
WHERE id = 6;

UPDATE academic_programmes SET
description = 'Analiza la producción y distribución de recursos. Economistas guían decisiones empresariales, diseñan políticas y fomentan el desarrollo económico, contribuyendo a sociedades prósperas y equitativas.'
WHERE id = 7;

UPDATE academic_programmes SET 
description = 'Dedicada al cuidado de la salud, la enfermería implica atender a pacientes, administrar tratamientos y promover el bienestar. Las enfermeras desempeñan un papel esencial en la asistencia médica, brindando cuidado directo y contribuyendo a la recuperación de las personas.'
WHERE id = 8;

UPDATE academic_programmes SET 
description = 'Indaga en cuestiones fundamentales, guiando el pensamiento crítico y ético. Su influencia radica en cultivar sociedades reflexivas y conscientes.'
WHERE id = 9;

UPDATE academic_programmes SET 
description = 'Estudia el pasado humano para comprender el presente y moldear el futuro. Historiadores analizan eventos, culturas y cambios sociales. Su impacto se refleja en la preservación de la identidad colectiva, la toma de decisiones informadas y la evolución continua de las sociedades.'
WHERE id = 10;

UPDATE academic_programmes SET 
description = 'Diseña y construye infraestructuras como puentes y carreteras. Ingenieros civiles mejoran entornos, garantizando seguridad y eficiencia. Su impacto es clave para el desarrollo urbano, conectividad y calidad de vida.'
WHERE id = 11;

UPDATE academic_programmes SET
description = 'Desarrolla y optimiza sistemas informáticos. Ingenieros de sistemas diseñan software, asegurando eficiencia y seguridad en la tecnología. Su impacto abarca desde innovaciones tecnológicas hasta la simplificación de procesos, contribuyendo al avance digital y la productividad.'
WHERE id = 12;

UPDATE academic_programmes SET 
description = 'Aplica principios químicos para diseñar procesos industriales. Ingenieros químicos crean productos y mejoran métodos, optimizando la producción de materiales esenciales. Su impacto se refleja en avances tecnológicos, sostenibilidad y contribución a diversas industrias.'
WHERE id = 13;

UPDATE academic_programmes SET
description = 'Aplica ciencia e ingeniería para producir alimentos seguros y eficientes. Innovan en procesos y aseguran calidad, contribuyendo a la producción de alimentos saludables.'
WHERE id = 14;

UPDATE academic_programmes SET
description = 'Analiza la sociedad y sus estructuras. Investigadores en ciencias sociales estudian comportamientos humanos, culturas y sistemas. Su impacto se refleja en la comprensión global, abordando desafíos sociales y contribuyendo a la construcción de comunidades más equitativas e informadas.'
WHERE id = 15;

UPDATE academic_programmes SET 
description = 'Se centra en el aprendizaje temprano de niños. Maestros de educación infantil cultivan habilidades básicas y emocionales. Su impacto es fundamental en el desarrollo inicial, preparando a los niños para la educación continua y construyendo bases sólidas para su futuro académico y social.'
WHERE id = 16;

UPDATE academic_programmes SET 
description = 'Explora expresiones artísticas a través de la palabra escrita. Escritores y críticos literarios analizan y crean obras. Su impacto es cultural, enriqueciendo la comprensión humana, transmitiendo valores y estimulando la imaginación, contribuyendo así a la diversidad y sofisticación cultural.'
WHERE id = 17;

UPDATE academic_programmes SET 
description = 'Estudia patrones, estructuras y relaciones cuantitativas. Matemáticos resuelven problemas y desarrollan teorías. Su impacto es fundamental en ciencia, tecnología y toma de decisiones. Contribuyen al avance del conocimiento y a la resolución de desafíos en diversas disciplinas.'
WHERE id = 18;

UPDATE academic_programmes SET 
description = 'Cuida la salud bucal. Odontólogos previenen y tratan problemas dentales. Su impacto es esencial para la salud general, mejorando sonrisas, previniendo enfermedades orales y contribuyendo al bienestar integral de las personas.'
WHERE id = 19;

UPDATE academic_programmes SET 
description = 'Explora la comunicación a través de diferentes lenguajes. Lingüistas analizan estructuras y evolución de idiomas. Su impacto es cultural y lingüístico, preservando la diversidad y facilitando la comprensión global, promoviendo así la comunicación y la conexión entre culturas.'
WHERE id = 20;

UPDATE academic_programmes SET 
description = 'Investiga la composición y propiedades de la materia. Químicos desarrollan nuevas sustancias y estudian reacciones. Su impacto es clave en medicina, industria y medio ambiente, contribuyendo al avance científico y al desarrollo de tecnologías sostenibles.'
WHERE id = 21;

UPDATE academic_programmes SET 
description = 'Se enfoca en el diseño y desarrollo de medicamentos. Químicos farmacéuticos crean compuestos para tratar enfermedades. Su impacto es crucial en la mejora de tratamientos médicos, contribuyendo a la salud y bienestar de la sociedad.'
WHERE id = 22;

UPDATE academic_programmes SET 
description = 'Aborda desafíos sociales. Trabajadores sociales apoyan a individuos y comunidades. Su impacto es vital para la equidad y el bienestar, interviniendo en situaciones de vulnerabilidad y promoviendo cambios positivos en la sociedad.'
WHERE id = 23;

UPDATE academic_programmes SET 
description = 'Se centra en la promoción y venta de productos. Profesionales de marketing crean estrategias para llegar a audiencias. Su impacto es crucial para el éxito comercial, conectando empresas con clientes y guiando decisiones de consumo en la sociedad.'
WHERE id = 24;

UPDATE academic_programmes SET 
description = 'Analiza y interpreta grandes conjuntos de datos. Científicos de datos extraen conocimientos y patrones. Su impacto es fundamental en la toma de decisiones informada, desde pronósticos comerciales hasta avances científicos, contribuyendo al progreso tecnológico y social.'
WHERE id = 25;

UPDATE academic_programmes SET 
description = 'Explora la mente y el comportamiento humano. Psicólogos estudian emociones y conductas, ofreciendo apoyo y tratamientos. Su impacto es vital para la salud mental, promoviendo el bienestar individual y contribuyendo a una sociedad más comprensiva y saludable.'
WHERE id = 26;

UPDATE academic_programmes SET 
description = 'Analiza estructuras gubernamentales y procesos políticos. Estudiosos en ciencias políticas investigan sistemas y formulación de políticas. Su impacto es crucial para entender y mejorar la gobernanza, promoviendo la participación cívica y contribuyendo al desarrollo democrático y la justicia social.'
WHERE id = 27;

UPDATE academic_programmes SET 
description = 'Se enfoca en transacciones comerciales globales. Profesionales en negocios internacionales gestionan operaciones entre países. Su impacto es fundamental para la economía global, fomentando relaciones comerciales y facilitando el intercambio de bienes y servicios a escala internacional.'
WHERE id = 28;

UPDATE academic_programmes SET 
description = 'Diseña y construye embarcaciones y estructuras marítimas. Ingenieros navales contribuyen al desarrollo de la industria naval, asegurando eficiencia y seguridad en la navegación. Su impacto es vital para el transporte marítimo, la exploración oceánica y la economía global.'
WHERE id = 29;

UPDATE academic_programmes SET 
description = 'Integra mecánica, electrónica y control. Diseña sistemas automatizados para innovación tecnológica y eficiencia en diversas industrias.'
WHERE id = 30;

UPDATE academic_programmes SET 
description = 'Se centra en el diseño y desarrollo de sistemas mecánicos. Ingenieros mecánicos crean máquinas y dispositivos. Su impacto es esencial en la industria, desde la fabricación hasta la energía, contribuyendo al progreso tecnológico y la resolución de desafíos ingenieriles.'
WHERE id = 31;

UPDATE academic_programmes SET 
description = 'Optimiza procesos y sistemas en empresas. Ingenieros industriales mejoran eficiencia y calidad. Su impacto es clave para la productividad empresarial, reducción de costos y mejora continua, contribuyendo al crecimiento económico y la competitividad.'
WHERE id = 32;

UPDATE academic_programmes SET 
description = 'Optimiza procesos y sistemas en empresas. Ingenieros industriales mejoran eficiencia y calidad. Su impacto es clave para la productividad empresarial, reducción de costos y mejora continua, contribuyendo al crecimiento económico y la competitividad.'
WHERE id = 33;

UPDATE academic_programmes SET 
description = 'Diseña sistemas eléctricos y electrónicos. Ingenieros eléctricos impulsan innovación en tecnologías, desde dispositivos electrónicos hasta infraestructuras de energía, contribuyendo al progreso de la sociedad.'
WHERE id = 34;

UPDATE academic_programmes SET 
description = 'Aplica ingeniería a la medicina. Ingenieros biomédicos desarrollan tecnologías para la salud, como dispositivos de diagnóstico y prótesis. Su impacto es esencial para avances médicos, mejorando la atención y calidad de vida.'
WHERE id = 35;

UPDATE academic_programmes SET 
description = '
Aborda desafíos medioambientales. Ingenieros ambientales diseñan soluciones sostenibles. Su impacto es vital para preservar recursos, gestionar residuos y mitigar impactos ambientales, contribuyendo al bienestar de la Tierra y la creación de comunidades más sostenibles.'
WHERE id = 36;

UPDATE academic_programmes SET 
description = 'Diseña espacios y estructuras. Arquitectos crean ambientes funcionales y estéticos. Su impacto es crucial en el entorno construido, desde viviendas hasta edificios icónicos, influyendo en la calidad de vida y el paisaje urbano.'
WHERE id = 37;

UPDATE academic_programmes SET 
description = 'Optimiza operaciones industriales. Ingenieros de procesos mejoran eficiencia y calidad en la producción. Su impacto es clave para la manufactura, garantizando procesos eficientes y sostenibles, contribuyendo al progreso industrial y económico.'
WHERE id = 38;

UPDATE academic_programmes SET 
description = 'Fomenta la exploración y disfrute de destinos. Profesionales en turismo gestionan experiencias de viaje. Su impacto es significativo en la economía local, preservación cultural y entendimiento global, promoviendo la diversidad y el intercambio cultural.'
WHERE id = 39;

UPDATE academic_programmes SET 
description = 'Investiga culturas y sociedades humanas. Antropólogos estudian costumbres y diversidad. Su impacto es crucial para comprender la evolución humana, abordar desafíos sociales y promover la apreciación de la diversidad cultural, contribuyendo a la construcción de sociedades más inclusivas.'
WHERE id = 40;

UPDATE academic_programmes SET 
description = 'Expresa creatividad a través de formas visuales. Artistas plásticos crean obras con diversos medios. Su impacto es estético y cultural, enriqueciendo la experiencia humana, transmitiendo ideas y contribuyendo a la diversidad artística en la sociedad.'
WHERE id = 41;

UPDATE academic_programmes SET 
description = 'Explora teorías y aplicaciones informáticas. Científicos de la computación desarrollan software, impulsan la innovación digital y contribuyen al avance tecnológico, desde algoritmos hasta inteligencia artificial.'
WHERE id = 42;

UPDATE academic_programmes SET 
description = 'Crea contenidos visuales para entretener y comunicar. Profesionales en cine y televisión producen narrativas visuales. Su impacto es cultural y emocional, moldeando la forma en que percibimos historias y generando conexiones globales a través del entretenimiento.'
WHERE id = 43;

UPDATE academic_programmes SET 
description = 'Comunica visualmente a través de imágenes. Diseñadores gráficos crean contenido atractivo. Su impacto es clave en la publicidad, medios y cultura visual, transmitiendo mensajes de manera impactante y creativa.'
WHERE id = 44;

UPDATE academic_programmes SET 
description = 'Crea productos funcionales y estéticos. Diseñadores industriales innovan en objetos cotidianos, mejorando su utilidad y aspecto, desde electrodomésticos hasta muebles, para contribuir al confort y la estética en la sociedad.'
WHERE id = 45;

UPDATE academic_programmes SET 
description = 'Explora idiomas y literatura. Filólogos analizan textos, dialectos y evolución lingüística, preservando culturas y enriqueciendo la comprensión intercultural.'
WHERE id = 46;

UPDATE academic_programmes SET 
description = 'Analiza datos para identificar patrones. Estadísticos interpretan información, crucial para decisiones informadas en ciencia y negocios, contribuyendo al conocimiento y eficacia.'
WHERE id = 47;

UPDATE academic_programmes SET 
description = 'Investiga las leyes del universo, abordando partículas y energía. Físicos impulsan avances científicos y tecnológicos, transformando nuestra comprensión del mundo y mejorando la calidad de vida.'
WHERE id = 48;

UPDATE academic_programmes SET 
description = 'Trata trastornos de habla y audición. Fonoaudiólogos mejoran la comunicación y deglución, impactando positivamente en la calidad de vida al abordar dificultades orales y auditivas.'
WHERE id = 49;

UPDATE academic_programmes SET 
description = 'Estudia la superficie terrestre, analizando climas, paisajes y poblaciones. Su impacto es clave para entender la distribución de recursos y desafíos socioeconómicos, contribuyendo al conocimiento global y a la planificación sostenible.'
WHERE id = 50;

UPDATE academic_programmes SET
description = 'Aplica conocimientos técnicos a la agricultura. Ingenieros agrícolas mejoran técnicas de cultivo y gestión de recursos, contribuyendo a la seguridad alimentaria y al desarrollo rural.'
WHERE id = 51;

UPDATE academic_programmes SET 
description = 'Optimiza procesos agrícolas. Ingenieros agrónomos mejoran cultivos y gestión de suelos, contribuyendo al desarrollo rural y la seguridad alimentaria.'
WHERE id = 52;

UPDATE academic_programmes SET
description = 'Cuida la salud animal. Veterinarios previenen y tratan enfermedades, asegurando bienestar y seguridad alimentaria, contribuyendo al equilibrio ecológico y la relación entre humanos y animales.'
WHERE id = 53;

UPDATE academic_programmes SET 
description = 'Expresa emociones a través del sonido. Músicos crean y ejecutan obras. Su impacto es cultural y emocional, enriqueciendo la vida, transmitiendo sentimientos y fomentando la conexión humana a través de la expresión artística.'
WHERE id = 54;

UPDATE academic_programmes SET 
description = 'Relaciona alimentos y salud. Nutricionistas promueven dietas equilibradas, impactando la salud y previniendo enfermedades.'
WHERE id = 55;

UPDATE academic_programmes SET 
description = 'Estudia la sociedad y sus interacciones. Sociólogos analizan patrones sociales, contribuyendo a entender desafíos y promover la equidad para construir comunidades justas y cohesionadas.'
WHERE id = 56;

UPDATE academic_programmes SET 
description = 'Optimiza producción animal. Zootecnistas gestionan cría y alimentación, vital para agricultura y ganadería. Contribuyen a seguridad alimentaria y desarrollo rural.'
WHERE id = 57;

UPDATE academic_programmes SET 
description = 'Aplica ingeniería a sistemas biológicos. Diseña soluciones en biotecnología y medio ambiente, impulsando avances en salud y sostenibilidad.'
WHERE id = 58;

UPDATE academic_programmes SET 
description = 'Erigiendo estructuras, constructores mejoran entornos urbanos, vital para funcionalidad y estética en la sociedad.'
WHERE id = 59;

UPDATE academic_programmes SET 
description = 'Optimiza procesos empresariales. Mejora eficiencia y toma de decisiones para el éxito y la competitividad, contribuyendo al crecimiento económico.'
WHERE id = 60;

UPDATE academic_programmes SET 
description = 'Automatiza sistemas. Ingenieros diseñan procesos automáticos, mejorando eficiencia industrial y contribuyendo a la automatización para elevar productividad y calidad.'
WHERE id = 61;

UPDATE academic_programmes SET 
description = '
Extracción responsable. Ingenieros gestionan procesos, asegurando eficiencia y sostenibilidad para un desarrollo económico responsable y la gestión adecuada de recursos.'
WHERE id = 62;

UPDATE academic_programmes SET 
description = 'Optimiza extracción y procesamiento. Ingenieros aseguran eficiencia en la industria energética, contribuyendo al suministro y desarrollo económico.'
WHERE id = 63;

UPDATE academic_programmes SET 
description = 'Aplica principios físicos. Ingenieros físicos diseñan tecnologías avanzadas. Su impacto es crucial para la innovación, desarrollando dispositivos desde electrónicos hasta médicos, contribuyendo al avance tecnológico y científico.'
WHERE id = 64;

UPDATE academic_programmes SET 
description = 'Sustentabilidad en bosques. Ingenieros forestales gestionan recursos naturales. Su impacto es vital para la conservación, manejo de bosques y biodiversidad, contribuyendo a la sostenibilidad ambiental y al equilibrio ecológico.'
WHERE id = 65;

UPDATE academic_programmes SET 
description = 'Analiza la tierra. Ingenieros evalúan terrenos y recursos para construcción, minería y gestión ambiental, contribuyendo a la seguridad estructural y la preservación del medio ambiente.'
WHERE id = 66;

UPDATE academic_programmes SET 
description = 'Cuida la salud visual. Optometristas evalúan y corrigen problemas oculares. Su impacto es esencial para la salud ocular, detectando y tratando condiciones visuales, contribuyendo al bienestar visual y la calidad de vida.'
WHERE id = 67;

UPDATE academic_programmes SET 
description = 'Recupera el movimiento. Fisioterapeutas tratan lesiones y rehabilitan. Esencial para la salud, restauran funcionalidad física y alivian dolores, mejorando calidad de vida y bienestar.'
WHERE id = 68;

UPDATE academic_programmes SET 
description = 'Expresa estilo a través del diseño. Diseñadores crean prendas y tendencias. Su impacto es cultural y estético, influyendo en la industria y en la forma en que nos expresamos visualmente.'
WHERE id = 69;

UPDATE academic_programmes SET 
description = 'Explora vida en los océanos. Biólogos marinos estudian ecosistemas acuáticos. Su impacto es crucial para la conservación marina, comprendiendo y preservando la biodiversidad, contribuyendo a la sostenibilidad y al conocimiento del océano.'
WHERE id = 70;

UPDATE academic_programmes SET 
description = 'Cuida la salud humana. Médicos diagnostican y tratan enfermedades. Esencial para el bienestar, contribuyen a la prevención y cura, mejorando la calidad de vida y extendiendo la esperanza de vida.'
WHERE id = 71;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
UPDATE academic_programmes SET description = '' WHERE id >= 1 AND id <= 71;
-- +goose StatementEnd
