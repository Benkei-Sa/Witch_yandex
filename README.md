# Witch_yandex
Тестовое задание от Яндекса на стажировку

A. Среднячком быть очень круто
Ограничение времени   1 секунда
Ограничение памяти   512Mb
Ввод   стандартный ввод или input.txt
Вывод   стандартный вывод или output.txt
В стране Блерляндия жили-были три ведьмы. Каждый год троица собиралась на шабаш, в течение которого они производили ритуал призыва демона МедиАнона.

МедиАнон отличался своеобразным нравом — если он видел, что волшебная сила какой-либо волшебницы строго больше сил двух других либо строго меньше двух других, то он начинал злиться и гоняться за этой ведьмой до самого рассвета.

Очевидно, что ни одна из трех ведьм не хочет, чтобы МедиАнон гонялся за ней. Поэтому каждая ведьма придумала (независимо от других ведьм) изменить свою волшебную силу так, чтобы МедиАнон за ней не погнался.

Так как ведьмы придумывали это независимо и не делились своими планами друг с другом, то каждая ведьма производит вычисления, исходя из предположения, что две другие ведьмы менять ничего не будут.

Ведьма может изменять свою волшебную силу, прибавляя или удаляя к ней целое число магиков (единица измерения волшебной силы). В то же время каждая ведьма не хочет слишком сильно нагружаться данным изменением, поэтому хочет изменить свою волшебную силу на минимально возможное количество магиков.

Соответственно, каждую ведьму интересует два значения:

    наименьшее (по абсолютному значению) количество магиков, на которые данная ведьма должна изменить свою волшебную силу, чтобы во время ритуала МедиАнон НЕ погнался за ней.
    за сколькими ведьмами будет гоняться МедиАнон, если данная ведьма проведет запланированное изменение своей волшебной силы.

Формат ввода
В единственной строке даны 3 целых числа ai (−109≤ai≤109), разделенные пробелами — текущие волшебные силы ведьм (в магиках).
Формат вывода
Выведите 3 строки. В i-й строке выведите два целых числа через пробел:

    наименьшее (по абсолютному значению) количество магиков, на которые i-я ведьма должна изменить свою волшебную силу ai, чтобы во время ритуала МедиАнон НЕ погнался за ней.
    за сколькими ведьмами будет гоняться МедиАнон, если i-я ведьма проведет запланированное изменение своей волшебной силы.

Напоминаем, что

    каждая ведьма лишь продумывает свой план, поэтому реальных изменений волшебной силы не производится;
    каждая ведьма считает себя умнее других, поэтому считает, что другие две ведьмы не будут изменять свою волшебную силу.

Пример 1
Ввод
Вывод

2 6 5

3 1
1 1
0 2

Пример 2
Ввод
Вывод

-3 0 -3

0 1
3 0
0 1

Примечания
В первом тесте силы ведьм равны соответственно [2,6,5].

    Если ведьма 1 увеличит свою волшебную силу на 3 магика (с 2 до 5), то её волшебная сила сравняется с волшебной силой ведьмы 3, поэтому МедиАнон не будет за ней гоняться. В то же время за 2-й ведьмой МедиАнон погонится (так как её волшебная сила 6 будет строго больше двух других сил 5 и 5).
    Если ведьма 2 уменьшит свою волшебную силу на 1 магик (с 6 до 5), то её волшебная сила сравняется с волшебной силой ведьмы 3, поэтому МедиАнон не будет за ней гоняться. В то же время за 1-й ведьмой МедиАнон погонится (так как её волшебная сила 2 будет строго меньше двух других сил 5 и 5).
    Ведьме 3 ничего не надо делать, ведь её волшебная сила и так уже не больше и не меньше, чем у других ведьм, поэтому МедиАнон не будет за ней гоняться. В то же время за 1-й и за 3-й ведьмами МедиАнон погонится (так как их волшебная сила 2 и 6 будут строго меньше и больше двух других).

Во втором тесте силы ведьм равны соответственно [−3,0,−3] (да, отрицательные волшебные силы тоже бывают).

Ведьме 1 ничего не надо делать, ведь её волшебная сила уже не строго меньше, чем у других ведьм, поэтому МедиАнон не будет за ней гоняться. В то же время за 2-й ведьмой МедиАнон погонится (так как её волшебная сила 0 будет строго больше двух других).
    Если ведьма 2 уменьшит свою волшебную силу на 3 магик (с 0 до −3), то её волшебная сила сравняется с волшебными силами ведьм 1 и 3, поэтому МедиАнон не будет за ней гоняться. После данного изменения у всех ведьм будет одинаковая волшебная сила, поэтому МедиАнон не будет ни за кем гоняться.
    Ведьме 3 ничего не надо делать, ведь её волшебная сила уже не строго меньше, чем у других ведьм, поэтому МедиАнон не будет за ней гоняться. В то же время за 2-й ведьмой МедиАнон погонится (так как её волшебная сила 0 будет строго больше двух других).